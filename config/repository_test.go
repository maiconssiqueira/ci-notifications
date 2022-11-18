package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/maiconssiqueira/ci-notifications/config"
)

func TestConfigRepository(t *testing.T) {
	t.Run("isValid", func(t *testing.T) {
		os.Setenv("GHTOKEN", "Testing")
		os.Setenv("ORGANIZATION", "Testing")
		os.Setenv("REPOSITORY", "Testing")

		var repo config.Repository
		repoConfig := repo.New()
		got := repoConfig.Github.Organization
		want := "Testing"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("isStruct", func(t *testing.T) {
		os.Setenv("GHTOKEN", "Testing")
		os.Setenv("ORGANIZATION", "Testing")
		os.Setenv("REPOSITORY", "Testing")
		var repo config.Repository
		got := repo.New()
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
