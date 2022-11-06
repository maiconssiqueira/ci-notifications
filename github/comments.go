package github

import (
	"encoding/json"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (g *Github) CommentInit(prNumber int, body string) *Github {
	config := config.New()
	return &Github{
		Organization: config.Github.Organization,
		Repository:   config.Github.Repository,
		Token:        config.Github.Token,
		Comments: comments{
			PrNumber: prNumber,
			Body:     body,
		},
	}
}

func (g *Github) SendComment(github *Github) (string, error) {
	payload, _ := json.Marshal(github.Comments)
	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/issues/" + strconv.Itoa(github.Comments.PrNumber) + "/comments")
	res := http.HttpPost(payload, url, "", github.Token)
	resPretty, _ := http.PrettyJson(res)

	return resPretty.String(), nil
}
