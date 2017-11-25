package cmd

import (
	"fmt"
	"github.com/mcordell/storyview/circle"
	"github.com/mcordell/storyview/config"
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

		jira.PrintIssue(issue)

		circleConfig, err = config.Circle()

		if err != nil {
			panic(err)
		}

		result, err := jira.GetIssuePullRequests(client, issue)

		if err != nil {
			panic(err)
		}

		printBranches(result.Branches)

		printPullRequests(result.PullRequests, &circleConfig)
	},
}

func printBranches(branches []*jira.Branch) {
	if len(branches) == 0 {
		return
	}

	spacer := "  "

	fmt.Printf("%sBranches:\n", spacer)

	for _, branch := range branches {
		fmt.Printf("%s%s%s\n", spacer, spacer, branch.Name)
	}
}

func printPullRequests(pullRequests []*jira.PullRequest, circleCredentials *config.CircleConfiguration) {
	if len(pullRequests) == 0 {
		return
	}

	spacer := "  "

	fmt.Printf("%sPull Requests:\n", spacer)

	for _, pullRequest := range pullRequests {
		fmt.Printf("%s%s%s\n", spacer, spacer, pullRequest.OneLine())
		circleClient := circle.BuildClient(circleCredentials)
		builds, err := circle.GetBuilds(circleClient, pullRequest.SourceBranch())
		if err == nil {
			for _, build := range builds {
				fmt.Printf("%s%s%s%s\n", spacer, spacer, spacer, circle.OneLineBuild(build))
			}
		}
	}
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}
