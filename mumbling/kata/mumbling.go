package kata

import "strings"

func Accum(s string) string {
	ss := []string{}
	for i := 0; i < len(s); i++ {
		s := strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
		ss = append(ss, s)
	}
	return strings.Join(ss, "-")
}
