package example

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type ExamplePlugin struct{}

const (
	NS = "oss example" // define Namespaces needed
	USAGE_PREFIX = "ibmcloud oss example " // Usage prefix. Usually is "ibmcloud <NS>"
)

// required. how to handle the command.
// args[0]: 	command name
// args[1:]: 	command arguments
func (p *ExamplePlugin) Run(context plugin.PluginContext, args []string) {
	// handle commands
}

// required. define namespaces of the command
func (p *ExamplePlugin) Namespaces() []plugin.Namespace {
	return []plugin.Namespace{
		// define namespaces needed
	}
}

// required. define plugin commands
func (p *ExamplePlugin) Commands() []plugin.Command {
	return []plugin.Command{
		// define commands needed
	}
}
