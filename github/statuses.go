package github

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/maiconssiqueira/ci-notifications/http"
)

var availableState = map[string]bool{
	"pending": true,
	"error":   true,
	"failure": true,
	"success": true,
}

var availableContext = map[string]bool{
	"ci/build":       true,
	"ci/deploy":      true,
	"ci/unittests":   true,
	"ci/codequality": true,
}

func (g *Github) statusesInit(context string, state string, description string, targetUrl string) {
	g.Statuses.Context = context
	g.Statuses.State = state
	g.Statuses.Description = description
	g.Statuses.TargetUrl = targetUrl
}

func Checks(context string, state string, description string, targetUrl string) (string, error) {
	if !availableState[state] || !availableContext[context] {
		err := fmt.Errorf(`
		This state or context reported [%v, %v] is invalid, it can be one of the following: 
		Available states:  [error, failure, pending, success] 
		Available contexts: [ci/build, ci/deploy, ci/unittests, ci/codequality]`, state, context)
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
