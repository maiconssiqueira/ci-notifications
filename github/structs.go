package github

import "os"

var (
	organization = os.Getenv("ORGANIZATION")
	repository   = os.Getenv("REPOSITORY")
	sha          = os.Getenv("SHA")
	bearer       = os.Getenv("GHTOKEN")
)

type status struct {
	Context     string `json:"context"`
	State       string `json:"state"`
	Description string `json:"description"`
	TargetUrl   string `json:"target_url"`
}

type releases struct {
	TagName              string `json:"tag_name"`
	TargetCommitish      string `json:"target_commitish"`
	Name                 string `json:"name"`
	Body                 string `json:"body"`
	Draft                bool   `json:"draft"`
	Prerelease           bool   `json:"prerelease"`
	GenerateReleaseNotes bool   `json:"generate_release_notes"`
}

type comments struct {
	PrNumber int    `json:"prNumber"`
	Body     string `json:"body"`
}

type Github struct {
	Organization string   `json:"organization"`
	Repository   string   `json:"repository"`
	Bearer       string   `json:"bearer"`
	Sha          string   `json:"sha"`
	Statuses     status   `json:"status"`
	Releases     releases `json:"releases"`
	Comments     comments `json:"comments"`
}
