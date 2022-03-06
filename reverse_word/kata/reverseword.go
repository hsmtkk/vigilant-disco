package kata

import "strings"

func ReverseWords(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
