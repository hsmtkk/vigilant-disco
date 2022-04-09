package findnthdigit

import (
	"log"
	"strconv"
)

func FindNthDigit(num, nth int) int {
	if nth <= 0 {
		return -1
	}
	if num < 0 {
		num = -num
	}
	digits := strconv.Itoa(num)
	index := len(digits) - nth
	if index < 0 {
		return 0
	} else {
		d := digits[index]
		n, err := strconv.Atoi(string(d))
		if err != nil {
			log.Fatal(err)
		}
		return n
	}
}
