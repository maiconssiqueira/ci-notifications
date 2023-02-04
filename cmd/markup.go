package cmd

import (
	"log"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/internal/http"
	"github.com/spf13/cobra"
)

var markupCmd = &cobra.Command{
	Use:   "markup",
	Short: "Set labels to PR",
	Long:  `Markup allow you to mark a pull request up with some label.`,
	RunE: func(_ *cobra.Command, _ []string) error {

		initMarkup := notify.InitMarkup(pullrequest, labels, *repoConf)

		var get http.Handler = &http.Contains{
			Method:      "GET",
			ContentType: "application/json",
			Token:       initMarkup.Token,
			Url:         (initMarkup.Url + "/issues/" + strconv.Itoa(initMarkup.Markup.Issue_number) + "/labels"),
		}
		var post http.Handler = &http.Contains{
			Method:      "POST",
			Content:     initMarkup.Markup.Markups,
			ContentType: "application/json",
			Token:       initMarkup.Token,
			Url:         (initMarkup.Url + "/issues/" + strconv.Itoa(initMarkup.Markup.Issue_number) + "/labels"),
		}
		res, err := notify.SendMarkup(initMarkup, &initMarkup.Markup, post, get)
		if err != nil {
			return err
		}
		log.Println(res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(markupCmd)
	markupCmd.Flags().IntVarP(&pullrequest, "pullrequest", "P", 0, `Pull Request number`)
	markupCmd.Flags().StringSliceVar(&labels, "labels", nil, `A string label to markup a pull request`)
	markupCmd.MarkFlagRequired("labels")
	markupCmd.MarkFlagRequired("pullrequest")
}
