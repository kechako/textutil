package textutil

import (
	"reflect"
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

var splitWordTests = []struct {
	in  string
	out []string
}{
	{"", nil},
	{"GolangTextUtil", []string{"Golang", "Text", "Util"}},
	{"GolangTEXTUtil", []string{"Golang", "TEXT", "Util"}},
	{"$$0123%%golang01234TEXT56789Util", []string{"$$", "0123", "%%", "golang", "01234", "TEXT", "56789", "Util"}},
	{"Golang  Text Util", []string{"Golang", "Text", "Util"}},
	{"  Golang TEXTUtil", []string{"Golang", "TEXT", "Util"}},
	{"$$0123 %%golang  01234TEXT 56789Util   ", []string{"$$", "0123", "%%", "golang", "01234", "TEXT", "56789", "Util"}},
}

func TestSplitWord(t *testing.T) {
	for _, tt := range splitWordTests {
		s := SplitWord(tt.in)
		if !reflect.DeepEqual(s, tt.out) {
			t.Errorf("SplitWord(%s) => %#v, want %#v", tt.in, s, tt.out)
		}
	}
}
