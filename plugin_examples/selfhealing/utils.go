package selfhealing

import (
	"fmt"
	"encoding/json"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
)

func isOutputJson(args []string) bool {
	for index, value := range args {
		if value == "--output" && args[index+1] == "json"{
			return true
		}
	}
	return false
}

func printError(msg string, outputJson bool) {
	if outputJson {
		fmt.Println("{}")
	} else {
		ui := terminal.NewStdUI()
		ui.Failed(msg)
	}
}

func panicError(err error, outputJson bool) {
	if outputJson {
		fmt.Println("{}")
	} else {
		panic(err)
	}
}

func handlePNPErrors(errorV PNPErrors, outputJson bool) {
	if outputJson {
		res, err := json.Marshal(errorV)
		if err != nil {
			fmt.Println("{}")
		} else {
			fmt.Println(string(res))
		}
	} else {
		ui := terminal.NewStdUI()
		ui.Failed(errorV.Errors[0].Detail)
	}
}