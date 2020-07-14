package selfhealing

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type SHPlugin struct{}

const (
	NS = "oss selfhealing"
	USAGE_PREFIX = "ibm oss selfhealing "
)

func (s *SHPlugin) Run(context plugin.PluginContext, args []string) {
	switch args[0] {
	case "token":
		s.handleToken(context, false)
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

func (s *SHPlugin) Namespaces() []plugin.Namespace {
	return []plugin.Namespace{
		plugin.Namespace{
			Name:        NS,
			Description: "selfhealing commands",
		},
	}
}

func (s *SHPlugin) Commands() []plugin.Command {
	return []plugin.Command{
		{
			Namespace: NS,
			Name:        "token",
			Description: "Get selfhealing token",
			Usage:       USAGE_PREFIX + "token",
		},
		{
			Namespace: NS,
			Name:        "processes",
			Description: "Get selfhealing processes in an environment",
			Usage:       USAGE_PREFIX + "processes -c ENV_CRN [-s START_DATE] [-e END_DATE] [--limit LIMIT] [--offset OFFSET] [--output json]",
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
					Description: "response data limit size, e.g. 25",
					HasValue: true,
				},
				{
					Name: "offset",
					Description: "response data offset, e.g. 50",
					HasValue: true,
				},
				{
					Name: "output",
					Description: "output format. support 'json' only",
					HasValue: true,
				},
			},
		},
		{
			Namespace: NS,
			Name:        "threads",
			Description: "Get threads of a selfhealing process",
			Usage:       USAGE_PREFIX + "threads -p PROCESS_ID",
			Flags: []plugin.Flag{
				{
					Name: "p",
					Description: "process ID",
					HasValue: true,
				},
				{
					Name: "output",
					Description: "output format. support 'json' only",
					HasValue: true,
				},
			},
		},
		{
			Namespace: NS,
			Name:        "script-log",
			Description: "Get script log for a thread in a selfhealing process",
			Usage:       USAGE_PREFIX + "script-log -p PROCESS_ID -t THREAD_ID",
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
				{
					Name: "output",
					Description: "output format. support 'json' only",
					HasValue: true,
				},
			},
		},
		{
			Namespace: NS,
			Name:        "stdout-log",
			Description: "Get stdout log for a thread in a selfhealing process",
			Usage:       USAGE_PREFIX + "stdout-log -p PROCESS_ID -t THREAD_ID",
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
				{
					Name: "output",
					Description: "output format. support 'json' only",
					HasValue: true,
				},
			},
		},
		{
			Namespace: NS,
			Name:        "stderr-log",
			Description: "Get stderr log for a thread in a selfhealing process",
			Usage:       USAGE_PREFIX + "stderr-log -p PROCESS_ID -t THREAD_ID",
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
				{
					Name: "output",
					Description: "output format. support 'json' only",
					HasValue: true,
				},
			},
		},
	}
}
