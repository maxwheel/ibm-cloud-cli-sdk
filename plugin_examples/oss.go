package main

import (
	// "fmt"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin_examples/blink"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin_examples/selfhealing"
)

type OSSPlugin struct{
	SubPlugins []OSSSubPlugin
	Namespaces []plugin.Namespace
	Commands []plugin.Command
}

type OSSSubPlugin interface{
	Namespaces() []plugin.Namespace
	Commands() []plugin.Command
	Run(plugin.PluginContext, []string)
}

func main() {
	op := new(OSSPlugin)
	// define subplugins. 
	// if new subplugin involved, create new one and push to OSSSubPlugin here
	subPlugins := []OSSSubPlugin{new(blink.BlinkPlugin), new(selfhealing.SHPlugin)}
	op.SubPlugins = subPlugins
	// default value
	op.Namespaces = []plugin.Namespace{
		plugin.Namespace{
			Name:        "oss",
		},
	}
	// set namespaces and commands
	for _, p := range subPlugins{
		for _, ns := range p.Namespaces(){
			op.Namespaces = append(op.Namespaces, ns)
		}
		for _, c := range p.Commands(){
			op.Commands = append(op.Commands, c)
		}
	}
	// start
	plugin.Start(op)
}

func (op *OSSPlugin) Run(context plugin.PluginContext, args []string) {
	namespace := context.CommandNamespace()
	// fmt.Println("looking for namespace:", namespace)
	// loop subPlugins and see if the namespace matches
	for _, p := range op.SubPlugins{
		for _, ns := range p.Namespaces(){
			if ns.Name == namespace{
				p.Run(context, args)
				// end search, break2
				goto endLoop
			}
		}
	}
	endLoop:
}

func (op *OSSPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "OSS-Tooling",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Namespaces: op.Namespaces,
		Commands: op.Commands,
	}
}
