package kata

import "strings"

func wave(words string) []string {
	results := []string{}
	for i := 0; i < len(words); i++ {
		upperd := upper(words, i)
		if upperd == words {
			continue
		}
		results = append(results, upperd)
	}
	return results
}

func upper(word string, index int) string {
	if index == 0 {
		return strings.ToUpper(string(word[index])) + word[index+1:]
	} else if index == len(word)-1 {
		return word[:index] + strings.ToUpper(string(word[index]))
	} else {
		return word[:index] + strings.ToUpper(string(word[index])) + word[index+1:]
	}
}
