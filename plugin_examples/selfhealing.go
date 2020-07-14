package main

import (
	"fmt"
	// "io/ioutil"
	// "strings"
	// "os"
	// "os/exec"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/common/rest"
)

type SHPlugin struct{}

func main() {
	plugin.Start(new(SHPlugin))
}

func (s *SHPlugin) Run(context plugin.PluginContext, args []string) {
	switch args[0] {
	case "token":
		s.handleToken(context)
	case "processes":
		s.getEnvProcesses(context, args[1:])
	case "threads":
		s.getProcessThreads(context, args[1:])
	case "script-log":
		s.getThreadLog(context, args[1:], "script_log")
	case "stdout-log":
		s.getThreadLog(context, args[1:], "stdout_log")
	case "stderr-log":
		s.getThreadLog(context, args[1:], "stderr_log")
	}
}

func (s *SHPlugin) handleToken(context plugin.PluginContext) string{
	token, err := context.RefreshIAMToken()
	if (err != nil) {
		panic(err)
	}
	ui := terminal.NewStdUI()
	ui.Say(terminal.PromptColor("New Self-healing token:"))
	ui.Say(token)
	return token
}

func (s *SHPlugin) getEnvProcesses(context plugin.PluginContext, args []string) {
	ui := terminal.NewStdUI()
	var crn, start, end string
	var limit, offset int
	var outputJson bool
	isValid := func(s string) bool {
		return s != "-c" && s != "-s" && s != "-e" && s != "--output" && s != "--limit" && s != "--offset"
	}
	for len(args) > 0 {
		if args[0] == "-c" {
			if len(args) > 1 && isValid(args[1]) {
				crn = args[1]
				args = args[2:]
			} else {
				ui.Failed("CRN not provided.")
				return
			}
		} else if args[0] == "-s" {
			if len(args) > 1 && isValid(args[1]) {
				start = args[1]
				args = args[2:]
			} else {
				ui.Failed("Start date/time empty")
				return
			}
		} else if args[0] == "-e" {
			if len(args) > 1 && isValid(args[1]) {
				end = args[1]
				args = args[2:]
			} else {
				ui.Failed("End date/time empty")
				return
			}
		} else if args[0] == "--limit" {
			if len(args) > 1 && isValid(args[1]) {
				limit, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				ui.Failed("Incorrect limit")
				return
			}
		} else if args[0] == "--offset" {
			if len(args) > 1 && isValid(args[1]) {
				offset, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				ui.Failed("Incorrect offset")
				return
			}
		} else if args[0] == "--output" {
			if len(args) > 1 && args[1] == "json" {
				outputJson = true
				args = args[2:]
			} else {
				ui.Failed("Output support 'json' only")
				return
			}
		} else {
			ui.Failed("Cannot recognize argument " + args[0])
			return
		}
	}
	// if (crn == "") {
	// 	ui.Failed("CRN not provided.")
	// 	return
	// }
	token, err := context.RefreshIAMToken()
	if (err != nil) {
		panic(err)
	}
	// ui.Say("crn: " + crn)
	// ui.Say("Start time: " + start)
	// ui.Say("End time: " + end)
	// rest call
	loadNext := true
	for loadNext {
		resp, successV, errorV, err := queryProcesses(token, crn, start, end, limit, offset)
		// fmt.Println(resp.StatusCode)
		if err != nil {
			loadNext = false
			panic(err)
		} else if resp.StatusCode < 200 || resp.StatusCode > 299{
			if outputJson {
				res, err := json.Marshal(errorV)
				if err != nil {
					// panic(err)
					var empty [0]string
					fmt.Println(empty)
				} else {
					fmt.Println(string(res))
				}
			} else {
				// format table
				ui.Failed(errorV.Errors[0].Detail)
			}
			loadNext = false
		} else {
			// 2xx OK
			if outputJson {
				res, err := json.Marshal(successV)
				if err != nil {
					// panic(err)
					var empty [0]string
					fmt.Println(empty)
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

type SHProcesses struct{
	First string `json:"first"`
	Href string `json:"href"`
	Last string `json:"last"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	next string `json:"offset"`
	Processes []SHProcess `json:"processes"`
	Total int `json:"total_count"`
}

type SHProcess struct{
	Href string `json:"href"`
	Id string `json:"id"`
	Status string `json:"status"`
	Threads HrefStruct `json:"threads"`
	Type string `json:"type"`
}

type HrefStruct struct{
	Href string `json:"href"`
}

type PNPErrors struct{
	Errors []PNPError `json:"errors"`
}

type PNPError struct{
	Code string `json:"code"`
	Detail string `json:"detail"`
	Message string `json:"message"`
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
	var outputJson bool
	isValid := func(s string) bool {
		return s != "-p" && s != "--limit" && s!= "--offset" && s != "--output"
	}
	for len(args) > 0 {
		if args[0] == "-p" {
			if len(args) > 1 && isValid(args[1]) {
				pid = args[1]
				args = args[2:]
			} else {
				ui.Failed("Process ID not provided.")
				return
			}
		} else if args[0] == "--limit" {
			if len(args) > 1 && isValid(args[1]) {
				limit, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				ui.Failed("Incorrect limit")
				return
			}
		} else if args[0] == "--offset" {
			if len(args) > 1 && isValid(args[1]) {
				offset, _ = strconv.Atoi(args[1])
				args = args[2:]
			} else {
				ui.Failed("Incorrect offset")
				return
			}
		} else if args[0] == "--output" {
			if len(args) > 1 && args[1] == "json" {
				outputJson = true
				args = args[2:]
			} else {
				ui.Failed("Output support 'json' only")
				return
			}
		} else {
			ui.Failed("Cannot recognize argument " + args[0])
			return
		}
	}
	token, err := context.RefreshIAMToken()
	if (err != nil) {
		panic(err)
	}
	// ui.Say("pid: " + pid)
	// rest call
	loadNext := true
	for loadNext {
		resp, successV, errorV, err := queryThreads(token, pid, limit, offset)
		// fmt.Println(resp.StatusCode)
		if err != nil {
			loadNext = false
			panic(err)
		} else if resp.StatusCode < 200 || resp.StatusCode > 299{
			if outputJson {
				res, err := json.Marshal(errorV)
				if err != nil {
					// panic(err)
					var empty [0]string
					fmt.Println(empty)
				} else {
					fmt.Println(string(res))
				}
			} else {
				// format table
				ui.Failed(errorV.Errors[0].Detail)
			}
			loadNext = false
		} else {
			// 2xx OK
			if outputJson {
				res, err := json.Marshal(successV)
				if err != nil {
					// panic(err)
					var empty [0]string
					fmt.Println(empty)
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

type SHThreads struct{
	First string `json:"first"`
	Href string `json:"href"`
	Last string `json:"last"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	next string `json:"offset"`
	Threads []SHThread `json:"processes"`
	Total int `json:"total_count"`
}

type SHThread struct{
	Href string `json:"href"`
	Id string `json:"id"`
	Status string `json:"status"`
	Logs HrefStruct `json:"logs"`
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
	ui := terminal.NewStdUI()
	var pid, tid string
	isValid := func(s string) bool {
		return s != "-p" && s != "-t"
	}
	for len(args) > 0 {
		if args[0] == "-p" {
			if len(args) > 1 && isValid(args[1]) {
				pid = args[1]
				args = args[2:]
			} else {
				ui.Failed("Process ID not provided.")
				return
			}
		} else if args[0] == "-t" {
			if len(args) > 1 && isValid(args[1]) {
				tid = args[1]
				args = args[2:]
			} else {
				ui.Failed("Thread ID not provided.")
				return
			}
		} else {
			ui.Failed("Cannot recognize argument " + args[0])
			return
		}
	}
	token, err := context.RefreshIAMToken()
	if (err != nil) {
		panic(err)
	}
	// ui.Say("pid: " + pid)
	// rest call
	url := fmt.Sprintf("https://pnp-api-oss.cloud.ibm.com/selfhealing/api/v1/processes/%s/threads/%s/logs/%s/content", pid, tid, logType)
	var r = rest.GetRequest(url)
	r.Add("Authorization", token)
	// rest client
	client := rest.NewClient()
	var successV struct{
		Content string `json:"content"`
	}
	// var successV interface{}
	var errorV interface{}
	resp, err := client.Do(r, &successV, &errorV)
	// fmt.Println("response: ", resp.StatusCode)
	if (err != nil) {
		panic(err)
	}
	if err != nil || resp.StatusCode < 200 || resp.StatusCode > 299 {
		res, err := json.Marshal(errorV)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res))
	} else {
		fmt.Println(successV.Content)
		// res, err := json.Marshal(successV)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(successV)
	}
}

func (s *SHPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "self-healing",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Namespaces: []plugin.Namespace{
			plugin.Namespace{
				Name:        "selfhealing",
			},
		},
		Commands: []plugin.Command{
			{
				Namespace: "selfhealing",
				Name:        "token",
				Description: "Get selfhealing token",
				Usage:       "ibmcloud selfhealing token",
			},
			{
				Namespace: "selfhealing",
				Name:        "processes",
				Description: "Get selfhealing processes",
				Usage:       "ibmcloud selfhealing processes -c ENV_CRN [-s START_DATE] [-e END_DATE] [--limit LIMIT] [--offset OFFSET] [--output json]",
				Flags: []plugin.Flag{
					{
						Name: "c",
						Description: "environment CRN",
						HasValue: true,
					},
					{
						Name: "s",
						Description: "start date/time, e.g.2020-06-21T10:30:00Z",
						HasValue: true,
					},
					{
						Name: "e",
						Description: "end date/time, e.g.2021-06-21T10:30:00Z",
						HasValue: true,
					},
					{
						Name: "limit",
						Description: "response data limit size, e.g.2021-06-21T10:30:00Z",
						HasValue: true,
					},
					{
						Name: "offset",
						Description: "response data offset, e.g.2021-06-21T10:30:00Z",
						HasValue: true,
					},
					{
						Name: "output",
						Description: "output format, support 'json' only",
						HasValue: true,
					},
				},
			},
			{
				Namespace: "selfhealing",
				Name:        "threads",
				Description: "Get threads of a selfhealing process",
				Usage:       "ibmcloud selfhealing threads -p PROCESS_ID",
				Flags: []plugin.Flag{
					{
						Name: "p",
						Description: "process ID",
						HasValue: true,
					},
				},
			},
			{
				Namespace: "selfhealing",
				Name:        "script-log",
				Description: "Get script log for a thread in a selfhealing process",
				Usage:       "ibmcloud selfhealing script-log -p PROCESS_ID -t THREAD_ID",
				Flags: []plugin.Flag{
					{
						Name: "p",
						Description: "process ID",
						HasValue: true,
					},
					{
						Name: "t",
						Description: "thread ID",
						HasValue: true,
					},
				},
			},
			{
				Namespace: "selfhealing",
				Name:        "stdout-log",
				Description: "Get stdout log for a thread in a selfhealing process",
				Usage:       "ibmcloud selfhealing stdout-log -p PROCESS_ID -t THREAD_ID",
				Flags: []plugin.Flag{
					{
						Name: "p",
						Description: "process ID",
						HasValue: true,
					},
					{
						Name: "t",
						Description: "thread ID",
						HasValue: true,
					},
				},
			},
			{
				Namespace: "selfhealing",
				Name:        "stderr-log",
				Description: "Get stderr log for a thread in a selfhealing process",
				Usage:       "ibmcloud selfhealing stderr-log -p PROCESS_ID -t THREAD_ID",
				Flags: []plugin.Flag{
					{
						Name: "p",
						Description: "process ID",
						HasValue: true,
					},
					{
						Name: "t",
						Description: "thread ID",
						HasValue: true,
					},
				},
			},
		},
	}
}
