package cmd

import (
	"fmt"

	"github.com/maiconssiqueira/ci-notifications/pkg/github/statuses"
	"github.com/spf13/cobra"
)

var statusesCmd = &cobra.Command{
	Use:   "statuses",
	Short: "Send updates to Github Checks",
	Long:  `Status checks allow you to send data related to tests or routines submitted to the repository via CI/CD pipelines.`,
	Run: func(cmd *cobra.Command, args []string) {

		res := statuses.GithubChecks(context, state, description, targetUrl)
		fmt.Println(res)
	},
}

var context string
var state string
var description string
var targetUrl string

func init() {
	rootCmd.AddCommand(statusesCmd)
	statusesCmd.Flags().StringVarP(&context, "context", "c", "", `A string label to differentiate this status from the status of other systems. 
	This field is case-insensitive`)
	statusesCmd.Flags().StringVarP(&targetUrl, "targetUrl", "t", "", `"The target URL to associate with this status. This URL will be linked from 
	the GitHub UI to allow users to easily see the source of the status. For example, if your continuous integration system is posting build status, 
	you would want to provide the deep link for the build output for this specific SHA: http://ci.example.com/user/repo/build/sha"`)
	statusesCmd.Flags().StringVarP(&state, "state", "s", "", "The state of the status. Can be one of: error, failure, pending, success")
	statusesCmd.Flags().StringVarP(&description, "description", "d", "", "The short description of the status")
	statusesCmd.MarkFlagRequired("context")
	statusesCmd.MarkFlagRequired("state")
	statusesCmd.MarkFlagRequired("description")
	statusesCmd.MarkFlagRequired("targetUrl")
}
