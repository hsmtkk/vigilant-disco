package freqseq

import (
	"strconv"
	"strings"
)

func FreqSeq(orig, separator string) string {
	charCounter := map[rune]int{}
	for _, ch := range orig {
		count, ok := charCounter[ch]
		if ok {
			count += 1
		} else {
			count = 1
		}
		charCounter[ch] = count
	}
	result := []string{}
	for _, ch := range orig {
		count := charCounter[ch]
		result = append(result, strconv.Itoa(count))
	}
	return strings.Join(result, separator)
}
