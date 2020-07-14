package blink

import (
	"fmt"
	"strings"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
)

func (b *BlinkPlugin) handleBlinkToken(context plugin.PluginContext) {
	token, err := context.RefreshIAMToken()
	if (err != nil) {
		panic(err)
	}
	ui := terminal.NewStdUI()
	ui.Say(terminal.PromptColor("New Blink token:"))
	// remove Bearer
	token = strings.Fields(token)[1]
	ui.Say(token)
	ui.Say("") // empty line

	// if set to env
	var email string
	var confirm bool
	err = ui.Prompt(terminal.PromptColor("Export blink http proxy in the terminal?"), nil).Resolve(&confirm)
	if (err != nil) {
		panic(err)
	}
	if (confirm) {
		// os.Setenv("blink_token", token)
		err = ui.Prompt("Email (IBMId)", &terminal.PromptOptions{}).Resolve(&email)
		if (err != nil) {
			panic(err)
		}
		ui.Say("Please copy and run the following commands:")
		ui.Say("")
		ui.Say(terminal.CommandColor(fmt.Sprintf("export http_proxy=\"http://%s:%s@blink.oss.cloud.ibm.com:1080\"", strings.Replace(email, "@", "%40", -1), token)))
		ui.Say("")
		ui.Say(terminal.CommandColor("export https_proxy=$http_proxy"))
		ui.Say("")
		ui.Ok()
	}
}
