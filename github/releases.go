package github

import (
	"bytes"
	"encoding/json"

	"github.com/maiconssiqueira/ci-notifications/http"
)

func (g *Github) releasesInit(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool) {
	g.Releases.TagName = tagName
	g.Releases.TargetCommitish = targetCommitish
	g.Releases.Name = name
	g.Releases.Body = body
	g.Releases.Draft = draft
	g.Releases.Prerelease = prerelease
	g.Releases.GenerateReleaseNotes = generateReleaseNotes
}

func Releases(tagName string, targetCommitish string, name string, body string, draft bool, prerelease bool, generateReleaseNotes bool) (string, error) {
	github := new(Github)
	github.releasesInit(tagName, targetCommitish, name, body, draft, prerelease, generateReleaseNotes)

	payload, _ := json.Marshal(github.Releases)
	url := ("https://api.github.com/repos/" + organization + "/" + repository + "/releases")
	res := http.HttpPost(payload, url, bearer)

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		panic(err)
	}

	return resPretty.String(), nil
}
