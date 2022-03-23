package kata

import "strings"

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}
	rows := []string{}
	for _, line := range strings.Split(s, "\n") {
		chars := []string{}
		for _, s := range line {
			chars = append(chars, strings.Repeat(string(s), k))
		}
		row := strings.Join(chars, "")
		for i := 0; i < n; i++ {
			rows = append(rows, row)
		}
	}
	return strings.Join(rows, "\n")
}
