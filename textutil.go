package textutil

import (
	"regexp"
	"strings"
	"unicode"
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

// SplitWord splits the text by word breaks.
func SplitWord(text string) []string {
	var words []string

	runes := []rune(text)
	offset := 0
	i := 0
	for ; i < len(runes); i++ {
		var prev, r, next rune
		r = runes[i]
		if i > 0 {
			prev = runes[i-1]
		}
		if i < len(runes)-1 {
			next = runes[i+1]
		}

		split, canAppend := isStartRune(r, prev, next)
		if i > offset && split {
			if canAppend {
				words = append(words, string(runes[offset:i]))
			}
			offset = i
		}
	}
	if offset < len(runes) {
		w := string(runes[offset:])
		if w != " " {
			words = append(words, w)
		}
	}

	return words
}

func isStartRune(r, prev, next rune) (result, canAppend bool) {
	switch {
	case unicode.IsUpper(r):
		if !unicode.IsUpper(prev) || unicode.IsLower(next) {
			result = true
		}
	case unicode.IsLower(r):
		if !unicode.IsUpper(prev) && !unicode.IsLower(prev) {
			result = true
		}
	case unicode.IsNumber(r):
		if !unicode.IsNumber(prev) {
			result = true
		}
	default:
		if unicode.IsSpace(prev) || unicode.IsUpper(prev) || unicode.IsLower(prev) || unicode.IsNumber(prev) {
			result = true
		}
	}
	if result && !unicode.IsSpace(prev) {
		canAppend = true
	}

	return
}
