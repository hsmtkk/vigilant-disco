package kata

import "unicode"

func Solve(s string) []int {
	upper := 0
	lower := 0
	number := 0
	special := 0
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			upper += 1
		} else if unicode.IsLower(ch) {
			lower += 1
		} else if unicode.IsDigit(ch) {
			number += 1
		} else {
			special += 1
		}
	}
	return []int{upper, lower, number, special}
}
