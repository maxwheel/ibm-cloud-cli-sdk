package blink

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type BlinkPlugin struct{}

const (
	NS_BLINK = "oss blink"
	USAGE_PREFIX = "ibmcloud oss blink "
)

func (p *BlinkPlugin) Run(context plugin.PluginContext, args []string) {
	switch args[0] {
	case "token":
		p.handleBlinkToken(context)
	}
}

func (p *BlinkPlugin) Namespaces() []plugin.Namespace {
	return []plugin.Namespace{
		plugin.Namespace{
			Name: NS_BLINK,
			Description: "Blink related commands",
		},
	}
}

func (p *BlinkPlugin) Commands() []plugin.Command {
	return []plugin.Command{
		{
			Namespace: NS_BLINK,
			Name:        "token",
			Description: "Get blink token",
			Usage:       USAGE_PREFIX + "token",
		},
	}
}
