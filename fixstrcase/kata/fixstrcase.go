package kata

import "strings"

func Solve(str string) string {
	lowers := countLowers(str)
	uppers := countUppers(str)
	if lowers >= uppers {
		return strings.ToLower(str)
	} else {
		return strings.ToUpper(str)
	}
}

func countLowers(str string) int {
	count := 0
	for _, ch := range str {
		if 'a' <= ch && ch <= 'z' {
			count += 1
		}
	}
	return count
}

func countUppers(str string) int {
	count := 0
	for _, ch := range str {
		if 'A' <= ch && ch <= 'Z' {
			count += 1
		}
	}
	return count
}
