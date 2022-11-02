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

func (g *Github) StatusesInit(sha string, context string, state string, description string, targetUrl string) *Github {
	config := config.New()
	return &Github{
		Organization: config.Github.Organization,
		Repository:   config.Github.Repository,
		Token:        config.Github.Token,
		Sha:          sha,
		Statuses: status{
			Context:     context,
			State:       state,
			Description: description,
			TargetUrl:   targetUrl,
		},
	}
}

func (g *Github) Checks(github *Github) (string, error) {
	if !availableState[github.Statuses.State] || !availableContext[github.Statuses.Context] {
		err := fmt.Errorf(`
		This state or context reported [%v, %v] is invalid, it can be one of the following: 
		Available states:  [error, failure, pending, success] 
		Available contexts: [ci/build, ci/deploy, ci/unittests, ci/codequality]`, github.Statuses.State, github.Statuses.Context)
		return "", err
	}

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
