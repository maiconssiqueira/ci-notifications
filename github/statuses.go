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

func (n *Notification) SendStatus(github *Github, callback Callbacks) (string, error) {
	var post http.HttpHandlers = &http.Post{
		Content:     github.Statuses,
		ContentType: "application/json",
		Token:       github.Token,
		Url:         (github.Url + "/statuses/" + github.Sha),
	}

	raw, _, err := post.HandlerPost()
	if err != nil {
		log.Fatal(err)
	}

	return callback.Response(raw, github)
}
