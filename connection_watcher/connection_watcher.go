package connection_watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
	filehelpers "github.com/turbot/go-kit/files"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe/cmdconfig"
	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/db/db_local"
	pb "github.com/turbot/steampipe/plugin_manager/grpc/proto"
	"github.com/turbot/steampipe/steampipeconfig"
	"github.com/turbot/steampipe/utils"
)

type ConnectionWatcher struct {
	fileWatcherErrorHandler   func(error)
	watcherError              error
	watcher                   *utils.FileWatcher
	onConnectionConfigChanged func(configMap map[string]*pb.ConnectionConfig)
	count                     int
}

func NewConnectionWatcher(onConnectionChanged func(configMap map[string]*pb.ConnectionConfig)) (*ConnectionWatcher, error) {
	w := &ConnectionWatcher{
		onConnectionConfigChanged: onConnectionChanged,
	}

	watcherOptions := &utils.WatcherOptions{
		Directories: []string{constants.ConfigDir()},
		Include:     filehelpers.InclusionsFromExtensions([]string{constants.ConfigExtension}),
		ListFlag:    filehelpers.FilesRecursive,

		OnChange: func(events []fsnotify.Event) {
			w.handleFileWatcherEvent(events)
		},
	}
	watcher, err := utils.NewWatcher(watcherOptions)
	if err != nil {
		return nil, err
	}
	w.watcher = watcher

	// set the file watcher error handler, which will get called when there are parsing errors
	// after a file watcher event
	w.fileWatcherErrorHandler = func(err error) {
		log.Printf("[WARN] Failed to reload connection config: %s", err.Error())
	}

	watcher.Start()

	log.Printf("[INFO] created ConnectionWatcher")
	return w, nil
}

func (w *ConnectionWatcher) handleFileWatcherEvent(e []fsnotify.Event) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[WARN] ConnectionWatcher caught a panic: %s", helpers.ToError(r).Error())
		}
	}()

	// ignore the first event - this is raised as soon as we start the watcher
	// (this is to avoid conflicting calls to refreshConnections between Steampipe and the watcher)
	w.count++
	if w.count == 1 {
		log.Printf("[TRACE] handleFileWatcherEvent ignoring first event")
		return
	}

	log.Printf("[TRACE] ConnectionWatcher handleFileWatcherEvent")
	config, err := steampipeconfig.LoadConnectionConfig()
	if err != nil {
		log.Printf("[WARN] Error loading updated connection config: %s", err.Error())
		return
	}
	log.Printf("[TRACE] loaded updated config")
	client, err := db_local.NewLocalClient(constants.InvokerConnectionWatcher)
	if err != nil {
		log.Printf("[WARN] Error creating client to handle updated connection config: %s", err.Error())
	}
	defer client.Close()

	log.Printf("[TRACE] loaded updated config")
	// set the global steampipe config
	steampipeconfig.GlobalConfig = config
	// update the viper default based on this loaded config
	cmdconfig.SetViperDefaults(config.ConfigMap())

	log.Printf("[TRACE] calling onConnectionConfigChanged")
	// convert config to format expected by plugin manager
	// (plugin manager cannot reference steampipe config to avoid circular deps)
	configMap := NewConnectionConfigMap(steampipeconfig.GlobalConfig.Connections)
	w.onConnectionConfigChanged(configMap)

	log.Printf("[TRACE] calling RefreshConnectionAndSearchPaths")

	// now refresh connections and search paths
	refreshResult := client.RefreshConnectionAndSearchPaths()
	if refreshResult.Error != nil {
		log.Printf("[WARN] error refreshing connections: %s", refreshResult.Error)
		return
	}

	// display any refresh warnings
	refreshResult.ShowWarnings()
}

func (w *ConnectionWatcher) Close() {
	w.watcher.Close()
}
