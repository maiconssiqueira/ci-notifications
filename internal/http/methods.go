package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func HttpPost(payload []byte, url string, contentType string, token string) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("Content-Type", contentType)

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

func PrettyJson(input []byte) (*bytes.Buffer, error) {
	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, input, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("pretty json %w", err)
	}
	return resPretty, nil
}
