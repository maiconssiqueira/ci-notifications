package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh"},
	Short:   "Send checks status, notifications and update repositories releases",
	Long:    "a simple way to send notifications, update status of repositories and releases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("It is necessary to declare at least one resource. Try [github statuses]")
	},
}

func init() {
	githubCmd.AddCommand(releasesCmd)
	githubCmd.AddCommand(statusesCmd)
	githubCmd.AddCommand(commentsCmd)
	rootCmd.AddCommand(githubCmd)
}
