package kata

import "strings"

func Compose(s1, s2 string) string {
	ss1 := strings.Split(s1, "\n")
	ss2 := strings.Split(s2, "\n")
	size := len(ss1)
	result := []string{}
	for i := 0; i < size; i++ {
		left := ss1[i][:i+1]
		right := ss2[size-i-1][:size-i]
		result = append(result, left+right)
	}
	return strings.Join(result, "\n")
}
