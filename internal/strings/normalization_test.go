package strings_test

import (
	"testing"

	nsstr "github.com/13k/night-stalker/internal/strings"
)

func TestSlugify(t *testing.T) {
	testCases := []struct {
		Subject  string
		Expected string
	}{
		{"   :)", ""},
		{"10Gu!.小十", "10gu-小十"},
		{"7Mad (smurf)", "7mad-smurf"},
		{"` AlaCrity -", "alacrity"},
		{"Raging-_-Potato", "raging-potato"},
		{"楚源Cy", "楚源cy"},
	}

	for testCaseIdx, testCase := range testCases {
		actual := nsstr.Slugify(testCase.Subject)

		if actual != testCase.Expected {
			t.Errorf("case %d: expected %q, got %q", testCaseIdx, testCase.Expected, actual)
		}
	}
}
