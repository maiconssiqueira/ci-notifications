package github

import (
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (n *notification) InitComment(prNumber int, body string, repo config.Repository) *github {
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

func (n *notification) SendComment(github *github) (string, error) {
	url := (github.Url + "/issues/" + strconv.Itoa(github.Comments.PrNumber) + "/comments")
	_, pretty, _ := http.Post(github.Comments, url, "", github.Token)
	return pretty, nil
}
