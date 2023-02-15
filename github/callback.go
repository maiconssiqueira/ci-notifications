package github

import (
	"encoding/json"
	"fmt"
	"os"

	"time"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// var log = logrus.New()

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

type Callbacks interface {
	Response(raw []byte, github *Github) (string, error)
}

type Return struct {
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type ReturnMarkup struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

func (g *status) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	if g.Return.Message != "" {
		return "", fmt.Errorf(fmt.Sprintf("Whoops, status of %v there was an error. %v", github.Url, g.Return.Message))
	}
	return fmt.Sprintf("hooray, status of %v was updated at %v", github.Url, g.Return.CreatedAt.String()), nil
}

func (g *markup) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("hooray, the label %v for %v was marked up", g.Markups.Labels[0], github.Url), nil
}

func (g *comments) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	if g.Return.Message != "" {
		return "", fmt.Errorf(fmt.Sprintf("Whoops, your comment for %v was not sent. %v", github.Url, g.Return.Message))
	}
	return fmt.Sprintf("hooray, the comment for %v was sent at %v", github.Url, g.Return.CreatedAt.String()), nil
}

func (g *releases) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	if g.Return.Message != "" {
		return "", fmt.Errorf(fmt.Sprintf("Whoops, this release %v for %v there was an error. %v ", g.TagName, github.Url, g.Return.Message))
	}
	return fmt.Sprintf("hooray, this release %v for %v was defined at %v", g.TagName, github.Url, g.Return.CreatedAt.String()), nil
}
