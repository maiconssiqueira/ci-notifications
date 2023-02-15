package cmd

import (
	"regexp"
	"strings"

	"github.com/maiconssiqueira/ci-notifications/internal/http"
	"github.com/maiconssiqueira/ci-notifications/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusesCmd = &cobra.Command{
	Use:   "statuses",
	Short: "Send updates to Github Checks",
	Long:  `Status checks allow you to send data related to tests or routines submitted to the repository via CI/CD pipelines.`,
	Run: func(_ *cobra.Command, _ []string) {
		type Config struct {
			States   map[string]bool `mapstructure:"states"`
			Contexts map[string]bool `mapstructure:"contexts"`
		}

		initStatuses := notify.InitStatuses(sha, context, state, description, targetUrl, *repoConf)

		var (
			conf Config
			post http.Handler = &http.Contains{
				Method:      "POST",
				Content:     initStatuses.Statuses,
				ContentType: "application/json",
				Token:       initStatuses.Token,
				Url:         (initStatuses.Url + "/statuses/" + initStatuses.Sha),
			}
		)

		viper.Unmarshal(&conf)

		if !conf.States[state] || !conf.Contexts[context] {
			log.Fatal(`this state or context reported [%v, %v] is invalid, it can be one of the following:
        available states:   [%v]
        available contexts: [%v]`, state, context, strings.Join(output.KeysByValue(conf.States, true), ", "), strings.Join(output.KeysByValue(conf.Contexts, true), ", "))
		}
		if len(sha) != 40 {
			log.Fatal(`please, check commit head SHA. SHA must be a 40 character SHA1`)
		}
		valid, _ := regexp.MatchString("^https://|http://", targetUrl)
		if !valid {
			log.Fatal(`please, check targetUrl. Target url must use http(s) scheme`)
		}

		res, err := notify.SendStatus(initStatuses, &initStatuses.Statuses, post)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(res)
	},
}

func init() {
	rootCmd.AddCommand(statusesCmd)
	statusesCmd.Flags().StringVarP(&sha, "sha", "S", "", `A commit head SHA`)
	statusesCmd.Flags().StringVarP(&context, "context", "c", "", `A string label to differentiate this status from the status of other systems. 
	This field is case-insensitive`)
	statusesCmd.Flags().StringVarP(&targetUrl, "targetUrl", "t", "", `The target URL to associate with this status. This URL will be linked from 
	the GitHub UI to allow users to easily see the source of the status. For example, if your continuous integration system is posting build status, 
	you would want to provide the deep link for the build output for this specific SHA: http://ci.example.com/user/repo/build/sha`)
	statusesCmd.Flags().StringVarP(&state, "state", "s", "", "The state of the status. Can be one of: error, failure, pending, success")
	statusesCmd.Flags().StringVarP(&description, "description", "d", "", "The short description of the status")
	statusesCmd.MarkFlagRequired("sha")
	statusesCmd.MarkFlagRequired("context")
	statusesCmd.MarkFlagRequired("state")
	statusesCmd.MarkFlagRequired("description")
	statusesCmd.MarkFlagRequired("targetUrl")
}
