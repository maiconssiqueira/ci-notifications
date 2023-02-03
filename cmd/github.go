package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh"},
	Short:   "Send checks status, notifications and update repositories releases",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("it is necessary to declare at least one resource")
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.AddCommand(releasesCmd)
	githubCmd.AddCommand(statusesCmd)
	githubCmd.AddCommand(commentsCmd)
	githubCmd.AddCommand(markupCmd)
}
