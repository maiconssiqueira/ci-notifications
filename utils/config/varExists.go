package config

import (
	"fmt"
	"os"
	"strings"
)

func VarExists(variables []string) error {
	withError := []string{}
	for _, variable := range variables {
		_, ok := os.LookupEnv(variable)
		if !ok {
			withError = append(withError, variable)
		}
	}
	if len(withError) > 1 {
		error := fmt.Errorf("some variables have not been defined. Check it out: %v", strings.Join(withError, ", "))
		return error
	}
	return nil
}
