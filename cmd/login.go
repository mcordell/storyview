package cmd

import (
	"bufio"
	"fmt"
	"github.com/mcordell/storyview/config"
	"github.com/spf13/cobra"
	"os"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store credentials in key ring for future login",
	Run: func(cmd *cobra.Command, args []string) {
		var r int
		for r != 1 {
			r = collectNextAction()()
		}
	},
}

// End quits the login actions
func End() int {
	fmt.Println("Done.")
	return 1
}

// StoreJIRACredentials collects the credentials for JIRA
func storeJIRACredentials() int {
	username := prompt("JIRA username")
	password := prompt("JIRA password")
	err := config.StoreJIRACredentials(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}

// StoreCircleCredentials
func storeCircleCredentials() int {
	username := prompt("Circle username")
	password := prompt("Circle API Token")
	err := config.StoreCircleCredentials(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}

func collectNextAction() func() int {
	for {
		fmt.Printf(`1) Store JIRA credentials
2) Store Circle credentials
q) Quit
`)
		action := prompt("Next Action:")

		switch action {
		case "1":
			return storeJIRACredentials
		case "2":
			return storeCircleCredentials
		case "q":
			fallthrough
		case "Q":
			return End
		default:
			fmt.Println("Invalid selection")
		}
	}
}

// prompt given a prompt collect the result
func prompt(p string) (result string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(p)
	result, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return result[:len(result)-1]
}

func init() {
	RootCmd.AddCommand(loginCmd)
}
