package simplestringreverse

import "strings"

func Reverse(s string) string {
	spaceIndexes := indexes(s)
	s = strings.ReplaceAll(s, " ", "")
	s = rev(s)
	for _, i := range spaceIndexes {
		s = s[:i] + " " + s[i:]
	}
	return s
}

func indexes(s string) []int {
	indexes := []int{}
	for i, ch := range s {
		if ch == ' ' {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func rev(s string) string {
	r := []rune(s)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	return string(res)
}
