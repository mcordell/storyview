package displays

import (
	"fmt"
	"github.com/mcordell/storyview/circle"
	"github.com/mcordell/storyview/config"
	"github.com/mcordell/storyview/jira"
)

type Displayer interface {
	Branches(branches []*jira.Branch)
	PullRequests(pullRequests []*jira.PullRequest, circleCredentials *config.CircleConfiguration)
	Issue(issue jira.Issue)
}

// Display ...
func Display(i jira.Issue, branches []*jira.Branch, prs []*jira.PullRequest, c *config.CircleConfiguration, d Displayer) {
	d.Issue(i)
	d.Branches(branches)
	d.PullRequests(prs, c)
}

type Plain struct {
}

func (p Plain) Issue(issue jira.Issue) {
	fmt.Printf("Issue %s: %s\n", issue.Issue.Key, issue.Issue.Fields.Summary)
}

func (p Plain) Branches(branches []*jira.Branch) {
	if len(branches) == 0 {
		return
	}

	spacer := "  "

	fmt.Printf("%sBranches:\n", spacer)

	for _, branch := range branches {
		fmt.Printf("%s%s%s\n", spacer, spacer, branch.Name)
	}
}

func (p Plain) PullRequests(pullRequests []*jira.PullRequest, circleCredentials *config.CircleConfiguration) {
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
