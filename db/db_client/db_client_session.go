package db_client

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/sethvargo/go-retry"
	"github.com/turbot/steampipe/db/db_common"
	"github.com/turbot/steampipe/utils"
)

func (c *DbClient) AcquireSession(ctx context.Context) *db_common.AcquireSessionResult {
	sessionResult := &db_common.AcquireSessionResult{}
	c.sessionInitWaitGroup.Add(1)
	defer c.sessionInitWaitGroup.Done()

	// get a database connection and query its backend pid
	// note - this will retry if the connection is bad
	databaseConnection, backendPid, err := c.getSessionWithRetries(ctx)
	if err != nil {
		sessionResult.Error = err
		return sessionResult
	}

	c.sessionsMutex.Lock()
	session, found := c.sessions[backendPid]
	if !found {
		session = db_common.NewDBSession(backendPid)
		session.LifeCycle.Add("created")
	}
	// we get a new *sql.Conn everytime. USE IT!
	session.Connection = databaseConnection
	sessionResult.Session = session
	c.sessionsMutex.Unlock()

	log.Printf("[TRACE] Got Session with PID: %d", backendPid)

	defer func() {
		// make sure that we close the acquired session, in case of error
		if sessionResult.Error != nil && databaseConnection != nil {
			databaseConnection.Close()
		}
	}()

	if c.ensureSessionFunc == nil {
		return sessionResult
	}

	if !session.Initialized {
		log.Printf("[TRACE] Session with PID: %d - waiting for init lock", backendPid)
		session.LifeCycle.Add("queued_for_init")

		err := c.parallelSessionInitLock.Acquire(ctx, 1)
		if err != nil {
			sessionResult.Error = err
			return sessionResult
		}
		c.sessionInitWaitGroup.Add(1)

		log.Printf("[TRACE] Session with PID: %d - waiting for init start", backendPid)
		session.LifeCycle.Add("init_start")
		err, warnings := c.ensureSessionFunc(ctx, session)
		session.LifeCycle.Add("init_finish")
		sessionResult.Warnings = warnings
		c.sessionInitWaitGroup.Done()
		c.parallelSessionInitLock.Release(1)
		if err != nil {
			sessionResult.Error = err
			return sessionResult
		}

		// if there is no error, mark session as initialized
		session.Initialized = true

		log.Printf("[TRACE] Session with PID: %d - init DONE", backendPid)
	}

	// update required session search path if needed
	if strings.Join(session.SearchPath, ",") != strings.Join(c.requiredSessionSearchPath, ",") {
		err := c.setSessionSearchPathToRequired(ctx, databaseConnection)
		if err != nil {
			sessionResult.Error = err
			return sessionResult
		}
		session.SearchPath = c.requiredSessionSearchPath
	}

	session.UpdateUsage()

	// now write back to the map
	c.sessionsMutex.Lock()
	c.sessions[backendPid] = session
	c.sessionsMutex.Unlock()

	log.Printf("[TRACE] Session with PID: %d - returning", backendPid)

	return sessionResult
}

func (c *DbClient) getSessionWithRetries(ctx context.Context) (*sql.Conn, int64, error) {
	backoff, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, 0, err
	}

	retries := 0
	var session *sql.Conn
	var backendPid int64
	const getSessionMaxRetries = 10
	err = retry.Do(ctx, retry.WithMaxRetries(getSessionMaxRetries, backoff), func(retryLocalCtx context.Context) (e error) {
		if utils.IsContextCancelled(retryLocalCtx) {
			return ctx.Err()
		}

		session, err = c.dbClient.Conn(retryLocalCtx)
		if err != nil {
			retries++
			return retry.RetryableError(err)
		}
		backendPid, err = db_common.GetBackendPid(retryLocalCtx, session)
		if err != nil {
			session.Close()
			retries++
			return retry.RetryableError(err)
		}
		return nil
	})

	if err != nil {
		log.Printf("[TRACE] getSessionWithRetries failed after 10 retries: %s", err)
		return nil, 0, err
	}

	if retries > 0 {
		log.Printf("[TRACE] getSessionWithRetries succeeded after %d retries", retries)
	}
	return session, backendPid, nil
}
