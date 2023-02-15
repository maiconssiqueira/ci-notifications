package config

import (
	"os"

	"github.com/maiconssiqueira/ci-notifications/internal/output"
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

type Repository struct {
	Github Github
}

type Github struct {
	Token        string
	Organization string
	Repository   string
	Url          string
}

func (r Repository) New() *Repository {
	err := output.CheckVariables("GHTOKEN", "ORGANIZATION", "REPOSITORY")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &Repository{
		Github: Github{
			Token:        os.Getenv("GHTOKEN"),
			Organization: os.Getenv("ORGANIZATION"),
			Repository:   os.Getenv("REPOSITORY"),
			Url:          "https://api.github.com/repos/" + os.Getenv("ORGANIZATION") + "/" + os.Getenv("REPOSITORY"),
		},
	}
}
