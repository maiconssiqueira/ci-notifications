package output

import (
	"fmt"
	"os"
	"strings"
)

func KeysByValue(m map[string]bool, value bool) []string {
	var keys []string
	for k, v := range m {
		if value == v {
			keys = append(keys, k)
		}
	}
	return keys
}

func CheckVariables(variables []string) error {
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
