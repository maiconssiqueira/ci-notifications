package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commentsCmd = &cobra.Command{
	Use:   "comments",
	Short: "The  comments supports send comment on pull requests",
	RunE: func(_ *cobra.Command, _ []string) error {
		init := gh.CommentInit(pullrequest, message, *repoConf)
		res, err := gh.SendComment(init)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var pullrequest int
var message string

func init() {
	rootCmd.AddCommand(commentsCmd)
	commentsCmd.Flags().IntVarP(&pullrequest, "pullrequest", "P", 0, `Pull Request number`)
	commentsCmd.Flags().StringVarP(&message, "message", "m", "", `Message to comment into a Pull Request`)
	commentsCmd.MarkFlagRequired("pullrequest")
	commentsCmd.MarkFlagRequired("message")
}
