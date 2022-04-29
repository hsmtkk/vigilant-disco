package kata

import "strings"

func Capitalize(st string, arr []int) string {
	result := []string{}
	for i, s := range st {
		t := string(s)
		if contains(i, arr) {
			t = strings.ToUpper(t)
		}
		result = append(result, t)
	}
	return strings.Join(result, "")
}

func contains(n int, ns []int) bool {
	for _, m := range ns {
		if n == m {
			return true
		}
	}
	return false
}
