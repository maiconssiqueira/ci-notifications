package github

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (g *Github) ReleasesInit(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool) *Github {
	config := config.New()
	return &Github{
		Organization: config.Github.Organization,
		Repository:   config.Github.Repository,
		Token:        config.Github.Token,
		Releases: releases{
			TagName:              tagName,
			TargetCommitish:      targetCommitish,
			Name:                 name,
			Body:                 body,
			Draft:                draft,
			Prerelease:           prerelease,
			GenerateReleaseNotes: generateReleaseNotes,
		},
	}
}

func (g *Github) SetRelease(github *Github) (string, error) {
	tagValidate, _ := regexp.MatchString("^(v[0-9]+)(\\.[0-9]+)(\\.[0-9])(|\\-rc\\.[0-9])(|\\-rc\\.[0-9])$", github.Releases.TagName)
	if !tagValidate {
		return "", fmt.Errorf(`this organization uses the semantic version pattern. You sent %v and the allowed is [v0.0.0, v0.0.0-rc0, v0.0.0-beta0]`, github.Releases.TagName)
	}

	payload, _ := json.Marshal(github.Releases)
	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/releases")
	res := http.HttpPost(payload, url, "application/json", github.Token)
	resPretty, _ := http.PrettyJson(res)

	return resPretty.String(), nil
}
