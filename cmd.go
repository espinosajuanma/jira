package jira

import (
	"os"

	"github.com/essentialkaos/go-jira/v2"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var Cmd = &Z.Cmd{
	Name:      `jira`,
	Summary:   `collection of jira helper commands`,
	Version:   `v0.0.1`,
	Commands:  []*Z.Cmd{help.Cmd, conf.Cmd, taskCmd},
	Shortcuts: Z.ArgMap{},
}

var taskCmd = &Z.Cmd{
	Name:     `task`,
	Summary:  `task-related commands`,
	Commands: []*Z.Cmd{help.Cmd, conf.Cmd, getTaskCmd},
}

var getTaskCmd = &Z.Cmd{
	Name:     `get`,
	Summary:  `print information of a particular task`,
	Usage:    `<id>`,
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		api, err := getJiraClient()
		if err != nil {
			return err
		}
		issue, err := api.GetIssue(
			args[0], jira.IssueParams{
				Expand: []string{"title"},
			},
		)
		if err != nil {
			return err
		}
		term.Printf("Assignee: %s\nCreator: %s\nSummary: %s\nDescription: %s\n", issue.Fields.Assignee.DisplayName, issue.Fields.Creator.DisplayName, issue.Fields.Summary, issue.Fields.Description)
		term.Printf("Decription: %s\n", issue.Fields.Description)
		return nil
	},
}

func getJiraClient() (*jira.API, error) {
	// FIXME: Use configuration files instead environment variables
	api, err := jira.NewAPI(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USER"), os.Getenv("JIRA_PASS"))
	api.SetUserAgent("Jira", "0.0.1")
	return api, err
}
