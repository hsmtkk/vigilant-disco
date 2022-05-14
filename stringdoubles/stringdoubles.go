package stringdoubles

import "strings"

func Shrink(orig string) string {
	stack := []string{}
	for _, ch := range orig {
		if len(stack) == 0 {
			stack = append(stack, string(ch))
		} else {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev != string(ch) {
				stack = append(stack, prev)
				stack = append(stack, string(ch))
			}
		}
	}
	return strings.Join(stack, "")
}
