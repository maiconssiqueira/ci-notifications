package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Github Github
}

type Github struct {
	Token        string
	Organization string
	Repository   string
}

func checkVariables(variables []string) error {
	withError := []string{}
	for _, variable := range variables {
		_, exists := os.LookupEnv(variable)
		if !exists || os.Getenv(variable) == "" {
			withError = append(withError, variable)
		}
	}
	if len(withError) > 0 {
		error := fmt.Errorf("some variables have not been defined or is empty. Check it out: %v", strings.Join(withError, ", "))
		return error
	}
	return nil
}

func New() *Config {
	checkVariables([]string{"GHTOKEN", "ORGANIZATION", "REPOSITORY"})
	return &Config{
		Github: Github{
			Token:        os.Getenv("GHTOKEN"),
			Organization: os.Getenv("ORGANIZATION"),
			Repository:   os.Getenv("REPOSITORY"),
		},
	}
}
