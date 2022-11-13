package output_test

import (
	"testing"

	"github.com/maiconssiqueira/ci-notifications/internal/output"
)

func TestOutput(t *testing.T) {
	t.Run("FindKeysByValues", func(t *testing.T) {
		toTest := map[string]bool{
			"Truthly": true,
			"Falsely": false,
		}
		want := []string{"Truthly"}
		got := output.KeysByValue(toTest, true)

		if got[0] != want[0] {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("EmptyKeysByValues", func(t *testing.T) {
		toTest := map[string]bool{
			"Truthly": false,
			"Falsely": false,
		}
		want := []string{}
		got := output.KeysByValue(toTest, true)

		if len(got) != 0 {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("InvalidSemVer", func(t *testing.T) {
		toTest := "v1.0.0-gopher1"
		want := "this organization uses the semantic version pattern. You sent " + toTest + " and the allowed is [v0.0.0, v0.0.0-rc0, v0.0.0-beta0]"
		got := output.CheckSemanticVersioning(toTest)

		if got.Error() != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("validSemVer", func(t *testing.T) {
		toTest := "v1.0.0-beta1"
		got := output.CheckSemanticVersioning(toTest)

		if got != nil {
			t.Errorf("got %v want %v", got, nil)
		}
	})
}
