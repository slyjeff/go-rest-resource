package restresource

import (
	"unicode"
	"unicode/utf8"
)

func makeCamelCase(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}

	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}

	return string(lc) + s[size:]
}
