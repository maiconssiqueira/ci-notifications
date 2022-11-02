package config

import (
	"fmt"
	"os"
	"strings"
)

var Vars = make(map[string]string)

func LoadVariables(variables []string) error {
	withError := []string{}
	for _, variable := range variables {
		_, ok := os.LookupEnv(variable)
		if !ok {
			withError = append(withError, variable)
		}
		Vars[variable] = os.Getenv(variable)
	}
	if len(withError) > 0 {
		error := fmt.Errorf("some variables have not been defined. Check it out: %v", strings.Join(withError, ", "))
		return error
	}
	return nil
}
