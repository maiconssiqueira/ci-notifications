package github

import (
	"encoding/json"
	"log"
	"time"
)

type Callbacks interface {
	Response(raw []byte, github *Github) (string, error)
}

type Return struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g *status) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	if g.Return.Message != "" {
		return "Whoops, status of " + github.Url + " there was an error. " + g.Return.Message, nil
	}
	return "Hooray, status of " + github.Url + " was updated at " + g.Return.CreatedAt.String(), nil
}

func (g *comments) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	if g.Return.Message != "" {
		return "Whoops, your comment for " + github.Url + " was not sent. " + g.Return.Message, nil
	}
	return "Hooray, the comment for " + github.Url + " was sent at " + g.Return.CreatedAt.String(), nil
}

func (g *releases) Response(raw []byte, github *Github) (string, error) {
	if err := json.Unmarshal(raw, &g.Return); err != nil {
		log.Fatal(err)
	}
	if g.Return.Message != "" {
		return "Whoops, this release " + g.TagName + " for " + github.Url + " there was an error. " + g.Return.Message, nil
	}
	return "Hooray, this release " + g.TagName + " for " + github.Url + " was defined at " + g.Return.CreatedAt.String(), nil
}
