package statuses

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/maiconssiqueira/ci-notifications/pkg/http"
)

type Status struct {
	Context     string `json:"context"`
	State       string `json:"state"`
	Description string `json:"description"`
	TargetUrl   string `json:"target_url"`
}

type Github struct {
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
	Bearer       string `json:"bearer"`
	Sha          string `json:"sha"`
	Statuses     Status `json:"status"`
}

func GithubChecks(context string, state string, description string, targetUrl string) string {

	github := Github{
		Organization: os.Getenv("ORGANIZATION"),
		Repository:   os.Getenv("REPOSITORY"),
		Sha:          os.Getenv("SHA"),
		Bearer:       os.Getenv("BEARER"),
		Statuses: Status{
			Context:     context,
			State:       state,
			Description: description,
			TargetUrl:   targetUrl,
		},
	}
	payload, _ := json.Marshal(github.Statuses)

	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/statuses/" + github.Sha)
	res := http.HttpPost(payload, url, github.Bearer)
	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		panic(err)
	}

	return resPretty.String()
}
