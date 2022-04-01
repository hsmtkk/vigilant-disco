package kata

import "strings"

func FindShort(s string) int {
	words := strings.Split(s, " ")
	min := len(words[0])
	for _, word := range words {
		if len(word) < min {
			min = len(word)
		}
	}
	return min
}
