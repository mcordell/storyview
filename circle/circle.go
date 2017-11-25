package circle

import (
	"fmt"
	"github.com/jszwedko/go-circleci"
	"github.com/mcordell/storyview/config"
	"github.com/mcordell/storyview/github"
	"github.com/pkg/errors"
)

// BuildClient creates a Circle CI client
func BuildClient(creds *config.CircleConfiguration) *circleci.Client {
	return &circleci.Client{Token: creds.Token}
}

// OneLineBuild returns a build as a single string line
func OneLineBuild(b *circleci.Build) string {
	return fmt.Sprintf("Num: %d, Status: %s, Outcome: %s | %s", b.BuildNum, b.Status, b.Outcome, b.BuildURL)
}

// GetBuilds gets the current builds from circle for a given branch
func GetBuilds(client *circleci.Client, b github.Branch) (builds []*circleci.Build, err error) {
	if b.Account == "" {
		err = errors.New("Cannot get builds for empty branch")
		return
	}

	return client.ListRecentBuildsForProject(
		b.Account,
		b.Repo,
		b.Branch,
		"",
		-1,
		0,
	)
}
