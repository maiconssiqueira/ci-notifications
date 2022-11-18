package github

import (
	"time"
)

//TODO
//IMAGINEI QUE SABIA, MAS NAO SEI USAR INTERFACES
// type Notifications interface {
// 	InitComment(prNumber int, body string, repo config.Repository) *Github
// 	SendComment(github *Github) (string, error)

// 	InitRelease(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool, repo config.Repository) *Github
// 	SetRelease(github *Github) (string, error)

//		InitStatuses(sha string, context string, state string, description string, targetUrl string, repo config.Repository) *Github
//		SendStatus(printLog bool, github *Github) (string, error)
//	}

type notification struct {
	Command []string
}

func NewNotification() *notification {
	return &notification{}
}

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

type github struct {
	Organization string `json:"organization"`
	Repository   string `json:"repository"`
	Url          string `json:"url"`
	Token        string
	Sha          string   `json:"sha"`
	Statuses     status   `json:"status"`
	Releases     releases `json:"releases"`
	Comments     comments `json:"comments"`
}
