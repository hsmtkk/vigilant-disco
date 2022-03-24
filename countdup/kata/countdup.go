package kata

import "strings"

func duplicate_count(s1 string) int {
	duplicated := map[rune]uint{}
	for _, ch := range strings.ToLower(s1) {
		count, ok := duplicated[ch]
		if ok {
			duplicated[ch] = count + 1
		} else {
			duplicated[ch] = 1
		}
	}
	result := 0
	for _, count := range duplicated {
		if count > 1 {
			result += 1
		}
	}
	return result
}
