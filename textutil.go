package textutil

import (
	"regexp"
	"strings"
)

var (
	undersocrePat1 = regexp.MustCompile(`([A-Z\d]+)([A-Z][a-z])`)
	undersocrePat2 = regexp.MustCompile(`([a-z\d]+)([A-Z])`)
	camelizePat1   = regexp.MustCompile(`^[a-z\d]*`)
	camelizePat2   = regexp.MustCompile(`^\w`)
	camelizePat3   = regexp.MustCompile(`[_-]([a-zA-Z\d]*)`)
)

// Underscore returns a copy of the string s with underscorizing.
func Underscore(s string) string {
	if m, _ := regexp.MatchString("[A-Z-]", s); !m {
		return s
	}

	s = undersocrePat1.ReplaceAllString(s, "${1}_${2}")
	s = undersocrePat2.ReplaceAllString(s, "${1}_${2}")
	s = strings.Replace(s, "-", "_", -1)
	s = strings.ToLower(s)

	return s
}

// Camelize returns a copy of the string s with camelizing
func Camelize(s string, upperFirst bool) string {
	if upperFirst {
		s = camelizePat1.ReplaceAllStringFunc(s, strings.Title)
	} else {
		s = camelizePat2.ReplaceAllStringFunc(s, strings.ToLower)
	}

	s = camelizePat3.ReplaceAllStringFunc(s, func(s string) string {
		return strings.Title(strings.ToLower(s[1:]))
	})

	return s
}
