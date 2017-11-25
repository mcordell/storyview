package displays

import (
	"fmt"
	"github.com/jszwedko/go-circleci"
	"github.com/mcordell/storyview/jira"
)

type StoryInformation struct {
	Issue        jira.Issue
	Branches     []*jira.Branch
	PRsAndBuilds PRsAndBuilds
}

type PRsAndBuilds map[*jira.PullRequest][]*circleci.Build

type Displayer interface {
	issue(jira.Issue)
	branches([]*jira.Branch)
	pullRequests(PRsAndBuilds)
}

// Display ...
func Display(info StoryInformation, displayer Displayer) {
	displayer.issue(info.Issue)
	displayer.branches(info.Branches)
	displayer.pullRequests(info.PRsAndBuilds)
}

// Plain is a plain text displayer
type Plain struct {
}

const plainSpacer = "  "

func (p Plain) issue(issue jira.Issue) {
	fmt.Printf("Issue %s: %s\n", issue.Issue.Key, issue.Issue.Fields.Summary)
	fmt.Printf("%s%s\n", plainSpacer, issue.BrowseURL)
}

func (p Plain) branches(branches []*jira.Branch) {
	if len(branches) == 0 {
		return
	}

	fmt.Printf("%sBranches:\n", plainSpacer)

	for _, branch := range branches {
		fmt.Printf("%s%s%s\n", plainSpacer, plainSpacer, branch.Name)
	}
}

func (p Plain) pullRequests(pnbs PRsAndBuilds) {
	if len(pnbs) == 0 {
		fmt.Println("PNBs are empty")
		return
	}

	fmt.Printf("%sPull Requests:\n", plainSpacer)

	for pullRequest, builds := range pnbs {
		fmt.Printf("%s%s%s\n", plainSpacer, plainSpacer, pullRequest.OneLine())
		for _, build := range builds {
			fmt.Printf("%s%s%s%s\n", plainSpacer, plainSpacer, plainSpacer, p.circleBuild(build))
		}
	}
}

func (p Plain) circleBuild(b *circleci.Build) string {
	return fmt.Sprintf("Num: %d, Status: %s, Outcome: %s | %s", b.BuildNum, b.Status, b.Outcome, b.BuildURL)
}

func (p Plain) pullRequest(pr jira.PullRequest) string {
	return fmt.Sprintf("%s %s", pr.Status, pr.URL)
}
