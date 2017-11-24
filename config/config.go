package config

import (
	"github.com/spf13/viper"
	"github.com/tmc/keyring"
)

const (
	// CircleUsernameConfigKey is configuration key for the circle username
	CircleUsernameConfigKey = "circleUsername"
	// JiraUsernameConfigKey is configuration key for the JIRA username
	JiraUsernameConfigKey = "jiraUsername"
	// JiraURLConfigKey is configuration key for the JIRA site url
	JiraURLConfigKey = "jiraURL"
)

// JIRAConfiguration is stores necessary data to create a JIRA client
type JIRAConfiguration struct {
	Username string
	Password string
	URL      string
}

// CircleConfiguration is stores necessary data to create a Circle client
type CircleConfiguration struct {
	Token string
}

// Circle returns a configuration for Circle API requests
func Circle() (c CircleConfiguration, err error) {
	var (
		username = viper.GetString(CircleUsernameConfigKey)
	)

	circleAPIToken, err := keyring.Get("hudCircle", username)
	if err != nil {
		return
	}

	return CircleConfiguration{Token: circleAPIToken}, nil
}

// JIRA returns the configuration for creating a JIRA client
func JIRA() (config JIRAConfiguration, err error) {
	var (
		username = viper.GetString(JiraUsernameConfigKey)
		url      = viper.GetString(JiraURLConfigKey)
	)

	config.Username = username
	config.URL = url

	jiraPassword, err := keyring.Get("storyviewJiraPass", username)

	if err != nil {
		return
	}

	config.Password = jiraPassword

	return
}
