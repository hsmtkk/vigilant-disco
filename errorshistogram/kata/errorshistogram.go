package kata

import (
	"fmt"
	"strings"
)

func Hist(s string) string {
	counter := CountChar(s)
	lines := []string{}
	for _, ch := range []rune{'u', 'w', 'x', 'z'} {
		count, ok := counter[ch]
		if ok {
			graph := strings.Repeat("*", int(count))
			line := fmt.Sprintf("%s  %d     %s", string(ch), count, graph)
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "\r")
}

func CountChar(s string) map[rune]uint {
	counter := map[rune]uint{}
	for _, ch := range s {
		if ch != 'u' && ch != 'w' && ch != 'x' && ch != 'z' {
			continue
		}
		count, ok := counter[ch]
		if ok {
			counter[ch] = count + 1
		} else {
			counter[ch] = 1
		}
	}
	return counter
}
