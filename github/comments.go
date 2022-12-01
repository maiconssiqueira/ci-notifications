package github

import (
	"log"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *Notification) InitComment(prNumber int, body string, repo config.Repository) *Github {
	return &Github{
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

func (n *Notification) SendComment(github *Github, callback Callbacks) (string, error) {
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
