package github

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/maiconssiqueira/ci-notifications/http"
)

func (g *Github) commentInit(prNumber int, body string) {
	g.Comments.PrNumber = prNumber
	g.Comments.Body = body
}

func Comment(prNumber int, body string) (string, error) {
	github := new(Github)
	github.commentInit(prNumber, body)
	payload, _ := json.Marshal(github.Comments)
	url := ("https://api.github.com/repos/" + organization + "/" + repository + "/issues/" + strconv.Itoa(prNumber) + "/comments")
	res := http.HttpPost(payload, url, "", bearer)

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		panic(err)
	}

	return resPretty.String(), nil
}
