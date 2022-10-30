package http

import (
	"bytes"
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
