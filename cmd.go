package jira

import (
	"fmt"

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
	Commands: []*Z.Cmd{help.Cmd, conf.Cmd, getTaskCmd, searchTaskCmd},
}

var getTaskCmd = &Z.Cmd{
	Name:     `get`,
	Summary:  `print information of a particular task`,
	Usage:    `<id>`,
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		api, err := getJiraClient(x.Caller.Caller)
		if err != nil {
			return err
		}
		prefix, _ := getIssuePrefix(x.Caller.Caller)
		issue, err := api.GetIssue(
			prefix+args[0], jira.IssueParams{
				Expand: []string{"title"},
			},
		)
		if err != nil {
			return err
		}
		term.Printf(`Assignee: %s
Creator: %s
Summary: %s
Description: %s`, issue.Fields.Assignee.DisplayName, issue.Fields.Creator.DisplayName, issue.Fields.Summary, issue.Fields.Description)
		return nil
	},
}

var searchTaskCmd = &Z.Cmd{
	Name:     `search`,
	Summary:  `find `,
	Usage:    `<status>`,
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, args ...string) error {
		api, err := getJiraClient(x.Caller.Caller)
		if err != nil {
			return err
		}
		status, _ := x.Caller.Caller.C("statuses." + args[0])
		if status == "null" {
			status = args[0]
		}
		res, _ := api.Search(jira.SearchParams{
			JQL: fmt.Sprintf("assignee=currentUser() and status='%s'", status),
		},
		)
		if res.Total == 0 {
			term.Print("There is no tasks")
			return nil
		}
		for _, issue := range res.Issues {
			term.Printf(`# %s - %s`, issue.Key, issue.Fields.Summary)
		}
		return nil
	},
}

func getJiraClient(cmd *Z.Cmd) (*jira.API, error) {
	url, _ := cmd.C("url")
	user, _ := cmd.C("user")
	pass, _ := cmd.C("pass")
	agent := cmd.Name
	version := cmd.Version
	api, err := jira.NewAPI(url, user, pass)
	api.SetUserAgent(agent, version)
	return api, err
}

func getIssuePrefix(cmd *Z.Cmd) (string, error) {
	return cmd.C("prefix")
}
