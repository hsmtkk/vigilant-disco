package kata

import (
	"strings"
)

func Solve(s string, a, b int) string {
	if b >= len(s) {
		b = len(s) - 1
	}
	left := s[:a]
	center := s[a : b+1]
	right := s[b+1:]
	return left + reverse(center) + right
}

func reverse(s string) string {
	reversed := []string{}
	for _, ch := range s {
		reversed = append([]string{string(ch)}, reversed...)
	}
	return strings.Join(reversed, "")
}
