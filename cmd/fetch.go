package cmd

import (
	"fmt"
	"github.com/mcordell/storyview/circle"
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

		pnbs := make(displays.PRsAndBuilds)

		circleClient := circle.BuildClient(&circleConfig)

		for _, pr := range result.PullRequests {
			builds, err := circle.GetBuilds(circleClient, pr.SourceBranch())
			pnbs[pr] = builds
			if err != nil && verbosity {
				fmt.Printf("Error during build fetching %s", err.Error())
			}
		}

		displays.Display(
			displays.StoryInformation{Issue: issue, Branches: result.Branches, PRsAndBuilds: pnbs},
			displays.Plain{},
		)
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}
