package doubleton

import "strconv"

func NextDoubleton(n int) int {
	for m := n + 1; ; m++ {
		if isDoubleton(m) {
			return m
		}
	}
}

func isDoubleton(n int) bool {
	digitCounter := map[rune]int{}
	for _, ch := range strconv.Itoa(n) {
		count, ok := digitCounter[ch]
		if ok {
			digitCounter[ch] = count + 1
		} else {
			digitCounter[ch] = 1
		}
	}
	return len(digitCounter) == 2
}
