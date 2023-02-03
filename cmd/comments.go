package cmd

import (
	"log"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/internal/http"
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
		initComment := notify.InitComment(pullrequest, message, *repoConf)

		var post http.Handler = &http.Contains{
			Method:      "POST",
			Content:     initComment.Comments,
			ContentType: "",
			Token:       initComment.Token,
			Url:         (initComment.Url + "/issues/" + strconv.Itoa(initComment.Comments.PrNumber) + "/comments"),
		}

		res, err := notify.SendComment(initComment, &initComment.Comments, post)
		if err != nil {
			return err
		}
		log.Println(res)
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
