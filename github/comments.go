package github

import (
	"log"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *Notification) InitComment(prNumber int, body string, repo config.Repository) *github {
	return &github{
		Organization: repo.Github.Organization,
		Repository:   repo.Github.Repository,
		Token:        repo.Github.Token,
		Url:          repo.Github.Url,
		Comments: comments{
			PrNumber: prNumber,
			Body:     body,
		},
	}
}

func (n *Notification) SendComment(github *github) (string, error) {
	var callback Callbacks = &github.Comments

	var post http.HttpHandlers = &http.Post{
		Content:     github.Comments,
		ContentType: "",
		Token:       github.Token,
		Url:         (github.Url + "/issues/" + strconv.Itoa(github.Comments.PrNumber) + "/comments"),
	}

	raw, _, err := post.HandlerPost()

	if err != nil {
		log.Fatal(err)
	}

	return callback.Response(raw, github)
}
