package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
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

func httpPost(payload []byte, url string, token string) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	defer req.Body.Close()

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	return data

}

func githubChecks(organization string, repository string, sha string, bearer string, context string, state string, description string, targetUrl string) string {
	github := Github{
		Organization: organization,
		Repository:   repository,
		Sha:          sha,
		Bearer:       bearer,
		Statuses: Status{
			Context:     context,
			State:       state,
			Description: description,
			TargetUrl:   targetUrl,
		},
	}

	payload, _ := json.Marshal(github.Statuses)
	url := ("https://api.github.com/repos/" + github.Organization + "/" + github.Repository + "/statuses/" + github.Sha)

	res := httpPost(payload, url, github.Bearer)

	return string(res)
}

func main() {
	res := githubChecks("FRST-Falconi",
		"calculadora-cpf-test",
		"54fca711aad922a1b286ca2415ebb0413cee86df",
		"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",
		"ci/pipelines", "success",
		"completed unit tests",
		"https://www.google.com",
	)

	fmt.Println(res)
}
