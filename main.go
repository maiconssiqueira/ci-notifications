package main

import (
	"fmt"
	"os"

	"github.com/maiconssiqueira/ci-notifications/pkg/github/statuses"
)

func main() {

	organization := os.Getenv("ORGANIZATION")
	repository := os.Getenv("REPOSITORY")
	sha := os.Getenv("SHA")
	bearer := os.Getenv("BEARER")
	context := os.Getenv("CONTEXT")
	state := os.Getenv("STATE")
	description := os.Getenv("DESCRIPTION")
	targetUrl := os.Getenv("TARGETURL")

	Github := statuses.Github{
		Organization: organization,
		Repository:   repository,
		Sha:          sha,
		Bearer:       bearer,
		Statuses: statuses.Status{
			Context:     context,
			State:       state,
			Description: description,
			TargetUrl:   targetUrl,
		},
	}
	res := statuses.GithubChecks(Github)

	fmt.Println(res)
}
