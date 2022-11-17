package github

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
)

func (g *Github) InitStatuses(sha string, context string, state string, description string, targetUrl string, repo config.Repository) *Github {
	return &Github{
		Organization: repo.Github.Organization,
		Repository:   repo.Github.Repository,
		Token:        repo.Github.Token,
		Url:          repo.Github.Url,
		Sha:          sha,
		Statuses: status{
			Context:     context,
			State:       state,
			Description: description,
			TargetUrl:   targetUrl,
		},
	}
}

func (g *Github) Checks(printLog bool, github *Github) (string, error) {
	url := (github.Url + "/statuses/" + github.Sha)
	raw, pretty, err := http.Post(github.Statuses, url, "application/json", github.Token)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &github.Statuses.ReturnStatuses); err != nil {
		log.Fatal(err)
	}

	if printLog {
		fmt.Println(pretty)
	}

	return "Hooray, status of " + github.Url + " was updated at " + github.Statuses.ReturnStatuses.CreatedAt.String(), nil
}
