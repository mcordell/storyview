package displays

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jszwedko/go-circleci"
	"github.com/mcordell/storyview/jira"
)

type Color struct {
}

func (c Color) issue(issue jira.Issue) {
	fmt.Printf("%s [%s] %s: %s\n", ColorizedType(issue.Issue.Fields.Type.Name), ColorizedStatus(issue.Issue.Fields.Status.Name), issue.Issue.Key, issue.Issue.Fields.Summary)
	fmt.Printf("%s%s\n", plainSpacer, issue.BrowseURL)
}

func (c Color) branches(branches []*jira.Branch) {
	if len(branches) == 0 {
		return
	}

	plainBold.Printf("%sBranches:\n", plainSpacer)

	for _, branch := range branches {
		fmt.Printf("%s%s%s\n", plainSpacer, plainSpacer, branch.Name)
	}
}

func (c Color) pullRequests(pnbs PRsAndBuilds) {
	if len(pnbs) == 0 {
		fmt.Println("PNBs are empty")
		return
	}

	plainBold.Printf("%sPull Requests:\n", plainSpacer)

	for pullRequest, builds := range pnbs {
		fmt.Printf("%s%s%s\n", plainSpacer, plainSpacer, c.pullRequest(*pullRequest))
		for _, build := range builds {
			fmt.Printf("%s%s%s%s\n", plainSpacer, plainSpacer, plainSpacer, c.circleBuild(build))
		}
	}
}

func (c Color) circleBuild(b *circleci.Build) string {
	return fmt.Sprintf("%s - %s ", b.BuildURL, ColorizedStatus(b.Outcome))
}

func (c Color) pullRequest(pr jira.PullRequest) string {
	return fmt.Sprintf("%s %s", ColorizedStatus(pr.Status), pr.URL)
}

// ColorizedStatus
func ColorizedStatus(s string) string {
	var c *color.Color
	c, ok := colorMap[s]
	if !ok {
		return s
	}
	return c.Sprint(s)
}

var greenText = color.New(color.FgGreen)
var yellowText = color.New(color.FgYellow)
var redText = color.New(color.FgRed)

var colorMap = map[string]*color.Color{
	"Task":        color.New(color.Bold, color.FgWhite, color.BgBlue),
	"Story":       color.New(color.Bold, color.FgWhite, color.BgGreen),
	"Bug":         color.New(color.Bold, color.FgWhite, color.BgRed),
	"success":     greenText,
	"Done":        greenText,
	"MERGED":      greenText,
	"CLOSED":      redText,
	"canceled":    redText,
	"OPEN":        yellowText,
	"In Progress": yellowText,
	"running":     yellowText,
}

var plainBold = color.New(color.Bold)

// ColorizedType ...
func ColorizedType(t string) string {
	var c *color.Color
	c, ok := colorMap[t]
	if !ok {
		c = plainBold
	}
	return c.Sprintf(" %s ", t)
}
