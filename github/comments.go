package github

import (
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

func (n *Notification) SendComment(github *Github) (string, error) {
	url := (github.Url + "/issues/" + strconv.Itoa(github.Comments.PrNumber) + "/comments")
	_, jsonPretty, _ := http.Post(github.Comments, url, "", github.Token)
	return jsonPretty, nil
}
