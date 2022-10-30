package main

import (
	"fmt"
)

func main() {
	res := githubChecks("FRST-Falconi", "calculadora-cpf-test", "54fca711aad922a1b286ca2415ebb0413cee86df", "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", "ci/pipelines", "success", "completed unit tests", "https://www.google.com")

	fmt.Println(res)
}
