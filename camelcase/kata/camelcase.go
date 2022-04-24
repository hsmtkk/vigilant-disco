package kata

import "strings"

func CamelCase(s string) string {
	results := []string{}
	for _, word := range strings.Split(s, " ") {
		results = append(results, CamelWord(word))
	}
	return strings.Join(results, "")
}

func CamelWord(s string) string {
	if s == "" {
		return ""
	}
	head := string(s[0])
	tail := s[1:]
	return strings.ToUpper(head) + tail
}
