package main

import (
	"log"

	"github.com/maiconssiqueira/ci-notifications/cmd"
	"github.com/maiconssiqueira/ci-notifications/utils/config"
)

var requiredVars = []string{
	"ORGANIZATION",
	"REPOSITORY",
	"SHA",
	"GHTOKEN",
}

func main() {
	err := config.LoadVariables(requiredVars)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()

}
