package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	wordSet := map[string]bool{}
	for _, x := range array1 {
		if contain(x, array2) {
			wordSet[x] = true
		}
	}
	result := []string{}
	for word := range wordSet {
		result = append(result, word)
	}
	sort.Strings(result)
	return result
}

func contain(s string, ss []string) bool {
	for _, t := range ss {
		if strings.Contains(t, s) {
			return true
		}
	}
	return false
}
