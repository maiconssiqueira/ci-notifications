package github

import (
	"encoding/json"
	"log"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *notification) InitStatuses(sha string, context string, state string, description string, targetUrl string, repo config.Repository) *github {
	return &github{
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

func (n *notification) SendStatus(github *github) (string, error) {
	url := (github.Url + "/statuses/" + github.Sha)
	raw, _, err := http.Post(github.Statuses, url, "application/json", github.Token)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &github.Statuses.ReturnStatuses); err != nil {
		log.Fatal(err)
	}

	if github.Statuses.ReturnStatuses.Message != "" {
		return "Whoops, status of " + github.Url + " there was an error. " + github.Statuses.ReturnStatuses.Message, nil
	}
	return "Hooray, status of " + github.Url + " was updated at " + github.Statuses.ReturnStatuses.CreatedAt.String(), nil
}
