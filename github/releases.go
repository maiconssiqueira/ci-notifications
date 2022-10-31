package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"regexp"

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
	tagValidate, _ := regexp.MatchString("^(v[0-9]+)(\\.[0-9]+)(\\.[0-9])(|\\-rc\\.[0-9])(|\\-rc\\.[0-9])$", tagName)
	if !tagValidate {
		err := fmt.Errorf(`
		This organization uses the semantic version pattern. You sent %v and the allowed is [v0.0.0, v0.0.0-rc0, v0.0.0-beta0]
		`, tagName)
		return "", err
	}
	github := new(Github)
	github.releasesInit(tagName, targetCommitish, name, body, draft, prerelease, generateReleaseNotes)

	payload, _ := json.Marshal(github.Releases)
	url := ("https://api.github.com/repos/" + organization + "/" + repository + "/releases")
	res := http.HttpPost(payload, url, "application/json", bearer)

	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return resPretty.String(), nil
}
