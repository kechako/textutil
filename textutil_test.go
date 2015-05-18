package textutil

import (
	"testing"
)

var underscoreTests = []struct {
	in  string
	out string
}{
	{"UnderscoreTestPattern", "underscore_test_pattern"},
	{"underscoreTestPattern", "underscore_test_pattern"},
	{"UNDERSCORE_TEST_PATTERN", "underscore_test_pattern"},
	{"UNDERSCORETestPATTERN", "underscore_test_pattern"},
	{"nounderscore", "nounderscore"},
}

func TestUnderscore(t *testing.T) {
	for _, tt := range underscoreTests {
		u := Underscore(tt.in)
		if u != tt.out {
			t.Errorf("Underscore(%s) => %s, want %s", tt.in, u, tt.out)
		}
	}
}

var camelizeTests = []struct {
	in         string
	upperFirst bool
	out        string
}{
	{"underscore_test_pattern", true, "UnderscoreTestPattern"},
	{"underscore_test_pattern", false, "underscoreTestPattern"},
}

func TestCamelize(t *testing.T) {
	for _, tt := range camelizeTests {
		u := Camelize(tt.in, tt.upperFirst)
		if u != tt.out {
			t.Errorf("Camelize(%s, %t) => %s, want %s", tt.in, tt.upperFirst, u, tt.out)
		}
	}
}
