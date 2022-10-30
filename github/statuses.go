package github

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/maiconssiqueira/ci-notifications/http"
)

var validState = map[string]bool{
	"pending": true,
	"error":   true,
	"failure": true,
	"success": true,
}

func (g *Github) statusesInit(context string, state string, description string, targetUrl string) {
	g.Statuses.Context = context
	g.Statuses.State = state
	g.Statuses.Description = description
	g.Statuses.TargetUrl = targetUrl
}

func Checks(context string, state string, description string, targetUrl string) (string, error) {
	if !validState[state] {
		err := fmt.Errorf("this state reported [%v] is invalid, it can be one of the following: error, failure, pending or success", state)
		return "", err
	}
	github := new(Github)
	github.statusesInit(context, state, description, targetUrl)

	payload, _ := json.Marshal(github.Statuses)
	url := ("https://api.github.com/repos/" + organization + "/" + repository + "/statuses/" + sha)
	res := http.HttpPost(payload, url, "application/json", bearer)

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		panic(err)
	}

	return resPretty.String(), nil
}
