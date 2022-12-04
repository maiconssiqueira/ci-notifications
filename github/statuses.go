package github

import (
	"log"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *Notification) InitStatuses(sha string, context string, state string, description string, targetUrl string, repo config.Repository) *Github {
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

func (n *Notification) SendStatus(github *Github, callback Callbacks, post http.HttpHandlers) (string, error) {
	raw, _, err := post.HandlerPost()
	if err != nil {
		log.Fatal(err)
	}

	return callback.Response(raw, github)
}
