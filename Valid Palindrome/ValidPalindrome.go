package main

import (
	"strings"
	"unicode"
)

func main() {

}

func validPalindrome(s string) bool {
	//применяем функцию для каждого символа строки и присваиваем результат исходной переменной
	s = strings.Map(isAlnum, s)

	l, r := 0, len(s)-1

	for l < r {
		// если символы не равны, возращаем false
		if s[l] != s[r] {
			return false
		}
		l++
		r++
	}
	return true
}

// проверяет что символ только буква или цифра
func isAlnum(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}
	return unicode.ToLower(r)
}
