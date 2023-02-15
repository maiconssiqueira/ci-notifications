package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = &logrus.Logger{
	Out:   os.Stderr,
	Level: logrus.DebugLevel,
	Formatter: &prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	},
}

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
