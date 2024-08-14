package two_pointers

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	pointerStart := 0
	pointerEnd := len(s) - 1

	for pointerStart < pointerEnd {
		if !unicode.IsLetter(rune(s[pointerStart])) && !unicode.IsNumber(rune(s[pointerStart])) {
			pointerStart += 1
			continue
		}

		if !unicode.IsLetter(rune(s[pointerEnd])) && !unicode.IsNumber(rune(s[pointerEnd])) {
			pointerEnd -= 1
			continue
		}

		if s[pointerStart] != s[pointerEnd] {
			return false
		}

		pointerStart += 1
		pointerEnd -= 1
	}

	return true
}
