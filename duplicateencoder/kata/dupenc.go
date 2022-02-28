package kata

import "strings"

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	charCounter := map[rune]uint{}
	for _, ch := range word {
		count, ok := charCounter[ch]
		if ok {
			charCounter[ch] = count + 1
		} else {
			charCounter[ch] = 1
		}
	}
	result := ""
	for _, ch := range word {
		count := charCounter[ch]
		if count > 1 {
			result += ")"
		} else {
			result += "("
		}
	}
	return result
}
