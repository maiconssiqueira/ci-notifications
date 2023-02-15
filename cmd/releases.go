package cmd

import (
	"github.com/maiconssiqueira/ci-notifications/internal/http"
	"github.com/maiconssiqueira/ci-notifications/internal/output"
	"github.com/spf13/cobra"
)

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Set a new release to a Github repository",
	Run: func(_ *cobra.Command, _ []string) {
		InitRelease := notify.InitRelease(tagName, targetCommitish, name, body, draft, prerelease, generateReleaseNotes, *repoConf)

		var post http.Handler = &http.Contains{
			Method:      "POST",
			Content:     InitRelease.Releases,
			ContentType: "application/json",
			Token:       InitRelease.Token,
			Url:         (InitRelease.Url + "/releases"),
		}

		err := output.CheckSemanticVersioning(tagName)
		if err != nil {
			log.Fatal(err)
		}
		res, err := notify.SetRelease(InitRelease, &InitRelease.Releases, post)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(res)
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)
	releasesCmd.Flags().StringVarP(&tagName, "tagName", "t", "", `The name of the tag. Example: v1.0.2`)
	releasesCmd.Flags().StringVarP(&targetCommitish, "targetCommitish", "T", "", `Specifies the commitish value that determines where the Git tag is created from. 
	Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository's default branch (usually master)`)
	releasesCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the release")
	releasesCmd.Flags().StringVarP(&body, "body", "b", "", "Text describing the contents of the tag. Markdown style")
	releasesCmd.Flags().BoolVarP(&draft, "draft", "d", false, "True to create a draft (unpublished) release, false to create a published one")
	releasesCmd.Flags().BoolVarP(&prerelease, "prerelease", "p", false, "True to identify the release as a prerelease. false to identify the release as a full release")
	releasesCmd.Flags().BoolVarP(&generateReleaseNotes, "generateReleaseNotes", "g", true, `Whether to automatically generate the name and body for this release. 
	If name is specified, the specified name will be used; otherwise, a name will be automatically generated. 
	If body is specified, the body will be pre-pended to the automatically generated notes`)
	releasesCmd.MarkFlagRequired("tagName")
	releasesCmd.MarkFlagRequired("targetCommitish")
	releasesCmd.MarkFlagRequired("name")
}
