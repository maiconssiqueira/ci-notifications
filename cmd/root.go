package cmd

import (
	"fmt"
	"os"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/github"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	tagName              string
	targetCommitish      string
	name                 string
	body                 string
	draft                bool
	prerelease           bool
	generateReleaseNotes bool
)

var (
	sha         string
	context     string
	state       string
	description string
	targetUrl   string
)

var (
	cfgFile  string
	repo     config.Repository
	notify   = github.NewNotification()
	repoConf = repo.New()
)

var (
	pullrequest int
	message     string
)

var (
	labels []string
)

var rootCmd = &cobra.Command{
	Use:   "ci-notifications",
	Short: "This is a simple and easy way to notify pipeline steps in your Github repository",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		_, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath("$PWD")
		viper.SetConfigName("config.yaml")
		viper.SetConfigType("yaml")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv() // read in environment variables that match

}
