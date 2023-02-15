package github

import (
	"encoding/json"
	"strings"

	"github.com/maiconssiqueira/ci-notifications/config"
	"github.com/maiconssiqueira/ci-notifications/internal/http"
	"github.com/maiconssiqueira/ci-notifications/internal/output"
)

func (n *Notification) InitMarkup(issue_number int, label []string, repo config.Repository) *Github {
	return &Github{
		Organization: repo.Github.Organization,
		Repository:   repo.Github.Repository,
		Token:        repo.Github.Token,
		Url:          repo.Github.Url,
		Markup: markup{
			Issue_number: issue_number,
			Markups: labels{
				label,
			},
		},
	}
}

func (n *Notification) SendMarkup(github *Github, callback Callbacks, post http.Handler, get http.Handler) (string, error) {
	raw, err := get.Request()
	if err != nil {
		log.Fatal(err)
	}

	var (
		response []ReturnMarkup
		found    []string
	)

	json.Unmarshal(raw, &response)
	for _, existent := range response {
		found = append(found, existent.Name)
	}

	github.Markup.Markups.Labels = output.CompareSlices(found, github.Markup.Markups.Labels)

	if len(found) > 0 && len(github.Markup.Markups.Labels) == 0 {
		log.Fatalf("whoops, These label(s) %v has already been marked up", strings.Join(found, ", "))
	} else if len(found) > 0 && len(github.Markup.Markups.Labels) > 0 {
		log.Infof("whoops, These label(s) %v has already been marked up", strings.Join(found, ", "))
	}

	raw, err = post.Request()
	if err != nil {
		log.Fatal(err)
	}

	return callback.Response(raw, github)
}
