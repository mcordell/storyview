package cmd

import (
	"github.com/mcordell/storyview/config"
	"github.com/mcordell/storyview/displays"
	"github.com/mcordell/storyview/jira"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch ISSUE",
	Short: "Fetch a JIRA issue",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			jiraConfig   config.JIRAConfiguration
			circleConfig config.CircleConfiguration
			err          error
		)

		jiraConfig, err = config.JIRA()

		if err != nil {
			panic(err)
		}

		client, err := jira.NewJIRAClient(&jiraConfig)

		if err != nil {
			panic(err)
		}

		issue, err := jira.GetIssue(client, args[0])

		if err != nil {
			panic(err)
		}

		circleConfig, err = config.Circle()

		if err != nil {
			panic(err)
		}

		result, err := jira.GetIssuePullRequests(client, issue.Issue)

		if err != nil {
			panic(err)
		}
		formatter := displays.Plain{}
		displays.Display(issue, result.Branches, result.PullRequests, &circleConfig, formatter)
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}
