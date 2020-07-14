package selfhealing

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/common/rest"
)

func (s *SHPlugin) handleToken(context plugin.PluginContext, noOutput bool) string{
	token, err := context.RefreshIAMToken()
	if (err != nil) {
		panic(err)
	}
	if noOutput == false {
		ui := terminal.NewStdUI()
		ui.Say(terminal.PromptColor("New Self-healing token:"))
		ui.Say(token)
	}
	return token
}

func (s *SHPlugin) getEnvProcesses(context plugin.PluginContext, args []string) {
	ui := terminal.NewStdUI()
	var crn, start, end string
	var limit, offset int
	outputJson := isOutputJson(args)
	isValid := func(s string) bool {
		return s != "-c" && s != "-s" && s != "-e" && s != "--output" && s != "--limit" && s != "--offset"
	}
	for len(args) > 0 {
		if args[0] == "-c" {
			if len(args) > 1 && isValid(args[1]) {
				crn = args[1]
				args = args[2:]
			} else {
				printError("CRN not provided.", outputJson)
				return
			}
		} else if args[0] == "-s" {
			if len(args) > 1 && isValid(args[1]) {
				start = args[1]
				args = args[2:]
			} else {
				printError("Start date/time empty", outputJson)
				return
			}
		} else if args[0] == "-e" {
			if len(args) > 1 && isValid(args[1]) {
				end = args[1]
				args = args[2:]
			} else {
				printError("End date/time empty", outputJson)
				return
			}
		} else if args[0] == "--limit" {
			if len(args) > 1 && isValid(args[1]) {
				limit, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				printError("Incorrect limit", outputJson)
				return
			}
		} else if args[0] == "--offset" {
			if len(args) > 1 && isValid(args[1]) {
				offset, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				printError("Incorrect offset", outputJson)
				return
			}
		} else {
			// printError("Cannot recognize argument " + args[0])
			// return
			// ignore unknown args
			args = args[1:]
		}
	}
	// check required params
	if crn == "" {
		printError("CRN not provided.", outputJson)
		return
	}
	// get token
	token := s.handleToken(context, true)
	// rest call
	loadNext := true
	for loadNext {
		resp, successV, errorV, err := queryProcesses(token, crn, start, end, limit, offset)
		// fmt.Println(resp.StatusCode)
		if err != nil {
			loadNext = false
			panicError(err, outputJson)
		} else if resp.StatusCode < 200 || resp.StatusCode > 299{
			handlePNPErrors(errorV, outputJson)
			loadNext = false
		} else {
			// 2xx OK
			if outputJson {
				res, err := json.Marshal(successV)
				if err != nil {
					fmt.Println("{}")
				} else {
					fmt.Println(string(res))
				}
				loadNext = false
			} else {
				lastIndex := offset + len(successV.Processes)
				ui.Say("")
				fmt.Printf("Processes: %d - %d (of total %d)\n", offset + 1, lastIndex, successV.Total)
				table := ui.Table([]string{"Process ID", "Status", "Type"})
				for _, value := range successV.Processes {
					table.Add(value.Id, value.Status, value.Type)
				}
				table.Print()
				if (lastIndex < successV.Total) {
					ui.Say("")
					err := ui.Prompt("Load more?", nil).Resolve(&loadNext)
					if err != nil {
						panic(err)
					}
					offset = lastIndex
				} else {
					loadNext = false
				}
			}
		}
	}
}

func queryProcesses(token string, crn string, start string, end string, limit int, offset int) (*http.Response, SHProcesses, PNPErrors, error) {
	url := "https://pnp-api-oss.cloud.ibm.com/selfhealing/api/v1/processes"
	var r = rest.GetRequest(url)
	r.Query("environment", crn)
	if start != "" {
		r.Query("start_time", start)
	}
	if end != "" {
		r.Query("end_time", end)
	}
	if limit != 0 {
		r.Query("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		r.Query("offset", strconv.Itoa(offset))
	}
	r.Add("Authorization", token)
	// rest client
	client := rest.NewClient()
	var successV SHProcesses
	var errorV PNPErrors
	resp, err := client.Do(r, &successV, &errorV)
	return resp, successV, errorV, err
}

func (s *SHPlugin) getProcessThreads(context plugin.PluginContext, args []string) {
	ui := terminal.NewStdUI()
	var pid string
	var limit, offset int
	outputJson := isOutputJson(args)
	isValid := func(s string) bool {
		return s != "-p" && s != "--limit" && s!= "--offset" && s != "--output"
	}
	for len(args) > 0 {
		if args[0] == "-p" {
			if len(args) > 1 && isValid(args[1]) {
				pid = args[1]
				args = args[2:]
			} else {
				printError("Process ID not provided.", outputJson)
				return
			}
		} else if args[0] == "--limit" {
			if len(args) > 1 && isValid(args[1]) {
				limit, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				printError("Incorrect limit", outputJson)
				return
			}
		} else if args[0] == "--offset" {
			if len(args) > 1 && isValid(args[1]) {
				offset, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				printError("Incorrect offset", outputJson)
				return
			}
		} else {
			// printError("Cannot recognize argument " + args[0], outputJson)
			// return
			// ignore unknown args
			args = args[1:]
		}
	}
	// check required params
	if pid == "" {
		printError("Process ID not provided.", outputJson)
		return
	}
	// get token
	token := s.handleToken(context, true)
	// rest call
	loadNext := true
	for loadNext {
		resp, successV, errorV, err := queryThreads(token, pid, limit, offset)
		// fmt.Println(resp.StatusCode)
		if err != nil {
			loadNext = false
			panicError(err, outputJson)
		} else if resp.StatusCode < 200 || resp.StatusCode > 299{
			handlePNPErrors(errorV, outputJson)
			loadNext = false
		} else {
			// 2xx OK
			if outputJson {
				res, err := json.Marshal(successV)
				if err != nil {
					fmt.Println("{}")
				} else {
					fmt.Println(string(res))
				}
				loadNext = false
			} else {
				lastIndex := offset + len(successV.Threads)
				firstIndex := offset + 1
				if firstIndex > successV.Total {
					firstIndex = successV.Total
				}
				ui.Say("")
				fmt.Printf("Threads : %d - %d (of total %d)\n", firstIndex, lastIndex, successV.Total)
				table := ui.Table([]string{"Thread ID", "Status"})
				for _, value := range successV.Threads {
					table.Add(value.Id, value.Status)
				}
				table.Print()
				if (lastIndex < successV.Total) {
					fmt.Println(lastIndex, successV.Total)
					ui.Say("")
					err := ui.Prompt("Load more?", nil).Resolve(&loadNext)
					if err != nil {
						panic(err)
					}
					offset = lastIndex
				} else {
					loadNext = false
				}
			}
		}
	}
}

func queryThreads(token string, pid string, limit int, offset int) (*http.Response, SHThreads, PNPErrors, error) {
	url := "https://pnp-api-oss.cloud.ibm.com/selfhealing/api/v1/processes/" + pid + "/threads"
	var r = rest.GetRequest(url)
	if limit != 0 {
		r.Query("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		r.Query("offset", strconv.Itoa(offset))
	}
	r.Add("Authorization", token)
	// rest client
	client := rest.NewClient()
	var successV SHThreads
	var errorV PNPErrors
	resp, err := client.Do(r, &successV, &errorV)
	return resp, successV, errorV, err
}

func (s *SHPlugin) getThreadLog(context plugin.PluginContext, args []string, logType string) {
	// ui := terminal.NewStdUI()
	var pid, tid string
	outputJson := isOutputJson(args)
	isValid := func(s string) bool {
		return s != "-p" && s != "-t"
	}
	for len(args) > 0 {
		if args[0] == "-p" {
			if len(args) > 1 && isValid(args[1]) {
				pid = args[1]
				args = args[2:]
			} else {
				printError("Process ID not provided.", outputJson)
				return
			}
		} else if args[0] == "-t" {
			if len(args) > 1 && isValid(args[1]) {
				tid = args[1]
				args = args[2:]
			} else {
				printError("Thread ID not provided.", outputJson)
				return
			}
		} else {
			// ui.Failed("Cannot recognize argument " + args[0])
			// return
			// ignore unkown args
			args = args[1:]
		}
	}
	// check required params
	if pid == "" {
		printError("Process ID not provided.", outputJson)
		return
	}
	if tid == "" {
		printError("Thread ID not provided.", outputJson)
		return
	}
	// get token
	token := s.handleToken(context, true)
	// rest call
	url := fmt.Sprintf("https://pnp-api-oss.cloud.ibm.com/selfhealing/api/v1/processes/%s/threads/%s/logs/%s/content", pid, tid, logType)
	var r = rest.GetRequest(url)
	r.Add("Authorization", token)
	// rest client
	client := rest.NewClient()
	var successV LogContent
	var errorV PNPErrors
	resp, err := client.Do(r, &successV, &errorV)
	// fmt.Println("response: ", resp.StatusCode)
	if (err != nil) {
		panicError(err, outputJson)
	} else if resp.StatusCode < 200 || resp.StatusCode > 299 {
		handlePNPErrors(errorV, outputJson)
	} else {
		if outputJson {
			res, err := json.Marshal(successV)
			if err != nil {
				panicError(err, outputJson)
			} else {
				fmt.Println(string(res))
			}
		} else {
			fmt.Println(successV.Content)
		}
	}
}
