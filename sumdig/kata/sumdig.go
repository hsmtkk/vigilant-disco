package kata

import (
	"log"
	"strconv"
)

func DigitalRoot(n int) int {
	ds := strconv.Itoa(n)
	if len(ds) == 1 {
		return n
	}
	sum := 0
	for _, d := range ds {
		m, err := strconv.Atoi(string(d))
		if err != nil {
			log.Fatal(err)
		}
		sum += m
	}
	return DigitalRoot(sum)
}
