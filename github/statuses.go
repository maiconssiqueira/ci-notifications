package github

import (
	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (g *Github) InitStatuses(sha string, context string, state string, description string, targetUrl string, repo config.Repository) *Github {
	return &Github{
		Organization: repo.Github.Organization,
		Repository:   repo.Github.Repository,
		Token:        repo.Github.Token,
		Url:          repo.Github.Url,
		Sha:          sha,
		Statuses: status{
			Context:     context,
			State:       state,
			Description: description,
			TargetUrl:   targetUrl,
		},
	}
}

func (g *Github) Checks(github *Github) (string, error) {
	url := (github.Url + "/statuses/" + github.Sha)
	json, _ := http.Post(github.Statuses, url, "application/json", github.Token)
	return json, nil
}
