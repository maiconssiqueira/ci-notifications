package config

import (
	"log"
	"os"

	"github.com/maiconssiqueira/ci-notifications/internal/output"
)

type Repository struct {
	Github github
}

type github struct {
	Token        string
	Organization string
	Repository   string
	Url          string
}

func (r Repository) New() *Repository {
	err := output.CheckVariables([]string{"GHTOKEN", "ORGANIZATION", "REPOSITORY"})
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{
		Github: github{
			Token:        os.Getenv("GHTOKEN"),
			Organization: os.Getenv("ORGANIZATION"),
			Repository:   os.Getenv("REPOSITORY"),
			Url:          "https://api.github.com/repos/" + os.Getenv("ORGANIZATION") + "/" + os.Getenv("REPOSITORY"),
		},
	}
}
