package config_test

import (
	"os"
	"testing"

	"github.com/maiconssiqueira/ci-notifications/utils/config"
)

func TestVars(t *testing.T) {
	want := ("some variables have not been defined. Check it out: MY1, MY2")
	var variableList = []string{
		"MY1",
		"MY2",
	}

	got := config.VarExists(variableList)
	if got.Error() != want {
		t.Errorf("Error actual = %v, and Expected = %v.", got, want)
	}
}

func TestValidVars(t *testing.T) {
	os.Setenv("MY1", "")
	os.Setenv("MY2", "")
	var variableList = []string{
		"MY1",
		"MY2",
	}

	got := config.VarExists(variableList)
	if got != nil {
		t.Errorf("Error actual = %v, and Expected = %v.", got, nil)
	}
}

func TestPArtialVars(t *testing.T) {
	want := ("some variables have not been defined. Check it out: MY2")
	os.Setenv("MY1", "")
	var variableList = []string{
		"MY1",
		"MY2",
	}

	got := config.VarExists(variableList)
	if got.Error() != want {
		t.Errorf("Error actual = %v, and Expected = %v.", got, want)
	}
}
