package main

import (
	"log"

	"github.com/maiconssiqueira/ci-notifications/cmd"
	"github.com/maiconssiqueira/ci-notifications/utils/config"
)

var variableList = []string{
	"ORGANIZATION",
	"REPOSITORY",
	"SHA",
	"GHTOKEN",
}

func main() {
	err := config.VarExists(variableList)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()

}
