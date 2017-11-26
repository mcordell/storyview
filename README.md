# Story view

> View JIRA issue details and associated branches, PRs, and builds in one command

![example image][example-image]

## Install

*With homebrew*

`brew install mcordell/tap/storyview`

or download the latest binaries [here][release-page]

## Setup

The below process only needs to be performed once.


Create a `.storyview.yaml` config file in your home directory:

```yaml
circleUsername: USER_NAME_TO_LOGIN_TO_CIRCLE
jiraUsername: USER_NAME_TO_LOGIN_TO_JIRA
jiraURL: JIRA_CLOUD_URL
```

Fill in your user names for circle and JIRA respectively. 'jiraURL' is the site
URL for your JIRA instance such as `https://your-corp.atlassian.net`.

### Authentication details

> All auth details / tokens are stored securely in macOS's keyring

Gather your [circle API token][circle-token-ref] and JIRA password.

From the command line, run:

```
storyview login
```

The command will prompt you for your credentials, type `q` when done.

## Usage

You can fetch the details of a story with one command `fetch`, just provide the
JIRA key after it as an argument.

```
storyview fetch PROJ-345
```

This will provide you with the following information:

![Annotated image][example-image-annotated]

*Protip* If using iTerm2, you can Cmd+Click on them to open them in browser

## Development

1. Install go
2. Install [dep][dep]
3. `git clone git@github.com:mcordell/storyview.git`
4. `dep ensure`


[dep]: https://github.com/golang/dep#setup
[example-image]: https://github.com/mcordell/storyview/raw/master/images/example.png
[example-image-annotated]: https://github.com/mcordell/storyview/raw/master/images/example-annotated.png
[circle-token-ref]: https://circleci.com/docs/api/v1-reference/#authentication
[release-page]: https://github.com/mcordell/storyview/releases/
