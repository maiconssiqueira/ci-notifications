package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/maiconssiqueira/ci-notifications/internal/output"
)

func httpPost(payload []byte, url string, contentType string, token string) []byte {
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
	data, _ := io.ReadAll(res.Body)
	return data
}

func Post(content any, url string, contentType string, token string) (string, error) {
	payload, _ := json.Marshal(content)
	res := httpPost(payload, url, contentType, token)
	resPretty, _ := output.PrettyJson(res)
	return resPretty.String(), nil
}
