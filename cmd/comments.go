package cmd

import (
	"fmt"

	"github.com/maiconssiqueira/ci-notifications/github"
	"github.com/spf13/cobra"
)

var (
	pullrequest int
	message     string
)

var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "The  comments supports send comment on pull requests",
	RunE: func(_ *cobra.Command, _ []string) error {
		var comments = &github.Github{}

		res, err := notify.SendComment(notify.InitComment(pullrequest, message, *repoConf), &comments.Comments)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(commentsCmd)
	commentsCmd.Flags().IntVarP(&pullrequest, "pullrequest", "P", 0, `Pull Request number`)
	commentsCmd.Flags().StringVarP(&message, "message", "m", "", `Message to comment into a Pull Request`)
	commentsCmd.MarkFlagRequired("pullrequest")
	commentsCmd.MarkFlagRequired("message")
}
