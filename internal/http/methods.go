package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func HttpHandler(method string, payload []byte, url string, contentType string, token string) []byte {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
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

func (h *Contains) Request() ([]byte, error) {
	payload, _ := json.Marshal(h.Content)
	res := HttpHandler(h.Method, payload, h.Url, h.ContentType, h.Token)
	return res, nil
}
