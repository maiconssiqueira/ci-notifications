package github

import (
	"encoding/json"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (g *Github) InitRelease(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool, repo config.Repository) *Github {
	return &Github{
		Organization: repo.Github.Organization,
		Repository:   repo.Github.Repository,
		Token:        repo.Github.Token,
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
	payload, _ := json.Marshal(github.Releases)
	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/releases")
	res := http.HttpPost(payload, url, "application/json", github.Token)
	resPretty, _ := http.PrettyJson(res)

	return resPretty.String(), nil
}
