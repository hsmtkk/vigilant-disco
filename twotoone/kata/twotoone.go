package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	joined := joinString(s1, s2)
	uniqued := unique(joined)
	sort.Strings(uniqued)
	return strings.Join(uniqued, "")
}

func joinString(s1, s2 string) []string {
	joined := []string{}
	for _, s := range s1 {
		joined = append(joined, string(s))
	}
	for _, s := range s2 {
		joined = append(joined, string(s))
	}
	return joined
}

func unique(ss []string) []string {
	m := map[string]bool{}
	for _, s := range ss {
		m[s] = true
	}
	uniqued := []string{}
	for k := range m {
		uniqued = append(uniqued, k)
	}
	return uniqued
}
