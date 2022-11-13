package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/maiconssiqueira/ci-notifications/config"
)

func TestConfigRepository(t *testing.T) {
	t.Run("missingVariables", func(t *testing.T) {
		os.Setenv("GHTOKEN", "Testing")

		var repo config.Repository
		repoConf, err := repo.New()
		got := err.Error()
		want := "some variables have not been defined or is empty. Check it out: ORGANIZATION, REPOSITORY"

		if got != want && repoConf == nil {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("isValid", func(t *testing.T) {
		os.Setenv("GHTOKEN", "Testing")
		os.Setenv("ORGANIZATION", "Testing")
		os.Setenv("REPOSITORY", "Testing")

		var repo config.Repository
		repoConfig, err := repo.New()
		got := repoConfig.Github.Organization
		want := "Testing"

		if got != want && err == nil {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("isStruct", func(t *testing.T) {
		os.Setenv("GHTOKEN", "Testing")
		os.Setenv("ORGANIZATION", "Testing")
		os.Setenv("REPOSITORY", "Testing")
		var repo config.Repository
		got, _ := repo.New()
		want := &config.Repository{
			Github: config.Github{
				Token:        "Testing",
				Organization: "Testing",
				Repository:   "Testing",
				Url:          "https://api.github.com/repos/Testing/Testing",
			},
		}

		if !(reflect.DeepEqual(got, want)) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
