package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/maiconssiqueira/ci-notifications/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

var statusesCmd = &cobra.Command{
	Use:   "statuses",
	Short: "Send updates to Github Checks",
	Long:  `Status checks allow you to send data related to tests or routines submitted to the repository via CI/CD pipelines.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		type Config struct {
			States   map[string]bool `mapstructure:"states"`
			Contexts map[string]bool `mapstructure:"contexts"`
		}

		var conf Config
		viper.Unmarshal(&conf)

		if !conf.States[state] || !conf.Contexts[context] {
			return fmt.Errorf(`this state or context reported [%v, %v] is invalid, it can be one of the following:
        available states:   [%v]
        available contexts: [%v]`, state, context, strings.Join(maps.Keys(conf.States), ", "), strings.Join(maps.Keys(conf.Contexts), ", "))
		}

		if len(sha) != 40 {
			return fmt.Errorf(`please, check commit head SHA. SHA must be a 40 character SHA1`)
		}

		valid, _ := regexp.MatchString("^https://|http://", targetUrl)
		if !valid {
			return fmt.Errorf(`please, check targetUrl. Target url must use http(s) scheme`)
		}

		github := github.Github{}
		init := github.StatusesInit(sha, context, state, description, targetUrl)
		res, err := github.Checks(init)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	},
}

var sha string
var context string
var state string
var description string
var targetUrl string

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
