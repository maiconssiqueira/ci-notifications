package github

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/http"
	"github.com/maiconssiqueira/ci-notifications/utils/config"
)

func (g *Github) commentInit(prNumber int, body string) {
	g.Comments.PrNumber = prNumber
	g.Comments.Body = body
}

func Comment(prNumber int, body string) (string, error) {
	github := new(Github)
	github.commentInit(prNumber, body)
	payload, _ := json.Marshal(github.Comments)
	url := ("https://api.github.com/repos/" + config.Organization + "/" + config.Repository + "/issues/" + strconv.Itoa(prNumber) + "/comments")
	res := http.HttpPost(payload, url, "", config.Bearer)

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return resPretty.String(), nil
}
