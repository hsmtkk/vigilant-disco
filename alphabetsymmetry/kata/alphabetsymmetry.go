package kata

import "strings"

const ALPHABETS = "abcdefghijklmnopqrstuvwxyz"

func solve(slice []string) []int {
	result := []int{}
	for _, s := range slice {
		result = append(result, count(s))
	}
	return result
}

func count(s string) int {
	s = strings.ToLower(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if ALPHABETS[i] == s[i] {
			count += 1
		}
	}
	return count
}
