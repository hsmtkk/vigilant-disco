package kata

import (
	"strconv"
	"strings"
)

func Numericals(s string) string {
	counter := map[rune]int{}
	result := []string{}
	for _, ch := range s {
		count := 1
		cnt, ok := counter[ch]
		if ok {
			count = cnt + 1
		}
		counter[ch] = count
		result = append(result, strconv.Itoa(count))
	}
	return strings.Join(result, "")
}
