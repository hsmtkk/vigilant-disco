package kata

import "strings"

func Solution(word string) string {
	result := []string{}
	for _, w := range word {
		result = append([]string{string(w)}, result...)
	}
	return strings.Join(result, "")
}
