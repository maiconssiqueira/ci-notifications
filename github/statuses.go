package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/maiconssiqueira/ci-notifications/http"
	"github.com/maiconssiqueira/ci-notifications/utils/config"
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
	g.Organization = config.Vars["ORGANIZATION"]
	g.Repository = config.Vars["REPOSITORY"]
	g.Token = config.Vars["GHTOKEN"]
	g.Sha = config.Vars["SHA"]

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
	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/statuses/" + github.Sha)
	res := http.HttpPost(payload, url, "application/json", github.Token)

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return resPretty.String(), nil
}
