package kata

import "strings"

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	words = CapitalizeFirstChar(words)
	return strings.Join(words, " ")
}

func CapitalizeFirstChar(words []string) []string {
	capitalized := []string{}
	for _, w := range words {
		capitalized = append(capitalized, strings.ToUpper(string(w[0]))+w[1:])
	}
	return capitalized
}
