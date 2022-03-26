package kata

import (
	"regexp"
	"strconv"
)

func SumOfIntegersInString(strng string) int {
	sum := 0
	for _, s := range Parse(strng) {
		n, err := strconv.Atoi(s)
		if err == nil {
			sum += n
		}
	}
	return sum
}

func Parse(s string) []string {
	result := []string{}
	pattern := regexp.MustCompile(`\d+`)
	for {
		loc := pattern.FindIndex([]byte(s))
		if loc == nil {
			break
		}
		begin := loc[0]
		end := loc[1]
		result = append(result, s[begin:end])
		s = s[end:]
	}
	return result
}
