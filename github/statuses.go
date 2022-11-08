package github

import (
	"encoding/json"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

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
	payload, _ := json.Marshal(github.Statuses)
	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/statuses/" + github.Sha)
	res := http.HttpPost(payload, url, "application/json", github.Token)
	resPretty, _ := http.PrettyJson(res)

	return resPretty.String(), nil
}
