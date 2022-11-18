package github

import (
	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *notification) InitRelease(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool, repo config.Repository) *github {
	return &github{
		Organization: repo.Github.Organization,
		Repository:   repo.Github.Repository,
		Token:        repo.Github.Token,
		Url:          repo.Github.Url,
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

func (n *notification) SetRelease(github *github) (string, error) {
	url := (github.Url + "/releases")
	_, jsonPretty, _ := http.Post(github.Releases, url, "application/json", github.Token)
	return jsonPretty, nil
}
