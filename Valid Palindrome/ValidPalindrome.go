package main

import (
	"strings"
	"unicode"
)

func main() {

}

func validPalindrome(s string) bool {
	s = strings.Map(isAlnum, s)

	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r++
	}
	return true
}

func isAlnum(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}
	return unicode.ToLower(r)
}
