package metaquery

import (
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/turbot/steampipe/autocomplete"
	"github.com/turbot/steampipe/schema"
	"github.com/turbot/steampipe/steampipeconfig"
)

// CompleterInput :: input interface for the metaquery completer
type CompleterInput struct {
	Query       string
	Schema      *schema.Metadata
	Connections *steampipeconfig.ConnectionDataMap
}

type completer func(input *CompleterInput) []prompt.Suggest

// Complete :: return completions for metaqueries.
func Complete(input *CompleterInput) []prompt.Suggest {
	input.Query = strings.TrimSuffix(input.Query, ";")
	cmd, _ := getCmdAndArgs(input.Query)

	metaQueryObj, found := metaQueryDefinitions[cmd]
	if !found {
		return []prompt.Suggest{}
	}
	if metaQueryObj.completer == nil {
		return []prompt.Suggest{}
	}
	return metaQueryObj.completer(input)
}

func completerFromArgsOf(cmd string) completer {
	return func(input *CompleterInput) []prompt.Suggest {
		metaQueryDefinition, _ := metaQueryDefinitions[cmd]
		suggestions := make([]prompt.Suggest, len(metaQueryDefinition.args))
		for idx, arg := range metaQueryDefinition.args {
			suggestions[idx] = prompt.Suggest{Text: arg.value, Description: arg.description, Output: arg.value}
		}
		return suggestions
	}
}

func inspectCompleter(input *CompleterInput) []prompt.Suggest {
	return autocomplete.GetTableAutoCompleteSuggestions(input.Schema, input.Connections)
}
