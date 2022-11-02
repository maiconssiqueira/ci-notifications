package github

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/http"
	"github.com/maiconssiqueira/ci-notifications/utils/config"
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

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return resPretty.String(), nil
}
