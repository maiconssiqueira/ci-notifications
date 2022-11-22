package github

import (
	"log"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *Notification) InitRelease(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool, repo config.Repository) *github {
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

func (n *Notification) SetRelease(github *github) (string, error) {
	var callback Callbacks = &github.Releases

	var post http.HttpHandlers = &http.Post{
		Content:     github.Releases,
		ContentType: "application/json",
		Token:       github.Token,
		Url:         (github.Url + "/releases"),
	}

	raw, _, err := post.HandlerPost()

	if err != nil {
		log.Fatal(err)
	}

	return callback.Response(raw, github)
}
