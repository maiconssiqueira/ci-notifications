package cmd

import (
	"fmt"

	"github.com/maiconssiqueira/ci-notifications/github"
	"github.com/spf13/cobra"
)

var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "The  comments supports send comment on pull requests",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := github.Comment(pullrequest, message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}

var pullrequest string
var message string

func init() {
	rootCmd.AddCommand(commentsCmd)
	commentsCmd.Flags().StringVarP(&pullrequest, "pullrequest", "P", "", `Pull Request number`)
	commentsCmd.Flags().StringVarP(&message, "message", "m", "", `Message to comment into a Pull Request`)
	commentsCmd.MarkFlagRequired("pullrequest")
	commentsCmd.MarkFlagRequired("message")
}
