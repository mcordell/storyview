package jira

import (
	"encoding/json"
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"github.com/mcordell/storyview/config"
	"github.com/pkg/errors"
	"io/ioutil"
	"regexp"
	"strings"
)

// NewJIRAClient creates a new JIRA client
func NewJIRAClient(creds *config.JIRAConfiguration) (c *jira.Client, err error) {
	c, err = jira.NewClient(nil, creds.URL)
	if err != nil {
		return
	}

	c.Authentication.SetBasicAuth(creds.Username, creds.Password)

	return
}

// GetIssue return a JIRA issue from by JIRA ID
func GetIssue(client *jira.Client, id string) (*jira.Issue, error) {
	issue, _, err := client.Issue.Get(id, nil)
	return issue, err
}

// PrintIssue ...
func PrintIssue(issue *jira.Issue) {
	fmt.Printf("Issue %s: %s\n", issue.Key, issue.Fields.Summary)
}

// GetIssuePullRequests gets Pull requests associated with a JIRA issue
func GetIssuePullRequests(client *jira.Client, issue *jira.Issue) (info GithubIssueInformation, err error) {
	baseURL := client.GetBaseURL()
	utwo := fmt.Sprintf(
		"%s/rest/dev-status/1.0/issue/detail?issueId=%s&applicationType=github&dataType=pullrequest",
		baseURL.String(),
		issue.ID,
	)
	r, err := client.NewRawRequest("GET", utwo, strings.NewReader(""))
	if err != nil {
		return
	}
	resp, err := client.Do(r, nil)

	if err != nil {
		result, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("%s", string(result))
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	response := GithubInfoResponse{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return
	}

	if len(response.Detail) != 0 {
		info = *response.Detail[0]
		return
	}

	err = errors.New("Response detail was empty")

	return
}

// GithubInfoResponse mirrors the response structure for a JIRA issue github
// info response request
type GithubInfoResponse struct {
	Detail []*GithubIssueInformation `json:"detail"`
}

// GithubIssueInformation contains the information for an Issue related to github items (branches/PRs)
type GithubIssueInformation struct {
	Branches     []*Branch      `json:"branches"`
	PullRequests []*PullRequest `json:"pullRequests"`
}

// Branch mirrors the data structure for a JIRA branch reference
type Branch struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GithubUser is a structure for representing a github user
type GithubUser struct {
	Name string `json:"name"`
}

// Reviewer is a structure for representing a github reviewer
type Reviewer struct {
	GithubUser
	Approved bool `json:"approved"`
}

// PullRequest is a structure for representing a github pull request
type PullRequest struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	CommentCount int         `json:"commentCount"`
	Author       *GithubUser `json:"author"`
	Status       string      `json:"status"`
	Reviewers    []*Reviewer `json:"reviewers"`
	URL          string      `json:"url"`
	Source       *Branch     `json:"source"`
}

// OneLineable is an interface for a structure that can be converted to a one line string
type OneLineable interface {
	OneLine() string
}

// OneLine returns a pull a request as a single string line
func (p *PullRequest) OneLine() string {
	return fmt.Sprintf("%s %s", p.Status, p.URL)
}

// SourceBranch returns the information for the source branch of a Pull Request
func (p *PullRequest) SourceBranch() GithubBranch {
	re := regexp.MustCompile("github.com/([^/]*)/([^/]*)/tree/([^/]*)")
	matches := re.FindAllSubmatch([]byte(p.Source.URL), -1)
	if len(matches) == 0 {
		return GithubBranch{}
	}

	var (
		account = string(matches[0][1])
		repo    = string(matches[0][2])
		branch  = string(matches[0][3])
	)

	return GithubBranch{
		Account: account,
		Repo:    repo,
		Branch:  branch,
	}
}

// GithubBranch stores basic data about a git branch on github.
type GithubBranch struct {
	Account string
	Repo    string
	Branch  string
}
