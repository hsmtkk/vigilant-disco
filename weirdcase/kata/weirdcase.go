package kata

import "strings"

func toWeirdCase(str string) string {
	words := strings.Split(str, " ")
	convertedWords := []string{}
	for _, word := range words {
		convertedWords = append(convertedWords, convert(word))
	}
	return strings.Join(convertedWords, " ")
}

func convert(word string) string {
	ss := []string{}
	for i, s := range word {
		if i%2 == 0 {
			ss = append(ss, strings.ToUpper(string(s)))
		} else {
			ss = append(ss, strings.ToLower(string(s)))
		}
	}
	return strings.Join(ss, "")
}
