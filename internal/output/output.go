package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	regexRelease = "^v([0-9]+)(\\.[0-9]+)(\\.[0-9]?)(|\\-rc)(|\\-beta)(|\\-alpha)([0-9])$"
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

func CheckVariables(variables ...string) error {
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

func CheckSemanticVersioning(semver string) error {
	valid, _ := regexp.MatchString(regexRelease, semver)
	if !valid {
		return fmt.Errorf(`this organization uses the semantic version pattern. You sent %v and the allowed is [v0.0.0, v0.0.0-rc0, v0.0.0-beta0]`, semver)
	}
	return nil
}

func PrettyJson(input []byte) (*bytes.Buffer, error) {
	resPretty := &bytes.Buffer{}
	err := json.Indent(resPretty, input, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("pretty json %w", err)
	}
	return resPretty, nil
}
