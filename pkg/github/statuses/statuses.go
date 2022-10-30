package statuses

import (
	"encoding/json"

	"github.com/maiconssiqueira/ci-notifications/pkg/http"
)

type Releases struct {
	TagName              string `json:"tag_name"`
	TargetCommitish      string `json:"target_commitish"`
	Name                 string `json:"name"`
	Body                 string `json:"body"`
	Draft                bool   `json:"draft"`
	Prerelease           bool   `json:"prerelease"`
	GenerateReleaseNotes bool   `json:"generate_release_notes"`
}

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

func GithubChecks(github Github) string {
	payload, _ := json.Marshal(github.Statuses)

	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/statuses/" + github.Sha)
	res := http.HttpPost(payload, url, github.Bearer)

	return string(res)
}
