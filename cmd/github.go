package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh"},
	Short:   "Send checks status, notifications and update repositories releases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("It is necessary to declare at least one resource. Try [github statuses]")
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.AddCommand(releasesCmd)
	githubCmd.AddCommand(statusesCmd)
	githubCmd.AddCommand(commentsCmd)
}
