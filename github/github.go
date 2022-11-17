package github

import "time"

type status struct {
	Context        string         `json:"context"`
	State          string         `json:"state"`
	Description    string         `json:"description"`
	TargetUrl      string         `json:"target_url"`
	ReturnStatuses returnStatuses `json:"return_status"`
}

type returnStatuses struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
	Url          string `json:"url"`
	Token        string
	Sha          string   `json:"sha"`
	Statuses     status   `json:"status"`
	Releases     releases `json:"releases"`
	Comments     comments `json:"comments"`
}
