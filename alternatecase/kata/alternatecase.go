package kata

import (
	"strings"
	"unicode"
)

func ToAlternatingCase(str string) string {
	result := ""
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			result += strings.ToLower(string(ch))
		} else if unicode.IsLower(ch) {
			result += strings.ToUpper(string(ch))
		} else {
			result += string(ch)
		}
	}
	return result
}
