package largestfive

import (
	"log"
	"strconv"
)

func Solve(digits string) uint {
	var max uint = 0
	for _, n := range Sequences(digits) {
		if n > max {
			max = n
		}
	}
	return max
}

func Sequences(digits string) []uint {
	sqs := []uint{}
	for i := 0; i <= len(digits)-5; i++ {
		ds := digits[i : i+5]
		n, err := strconv.Atoi(ds)
		if err != nil {
			log.Fatalf("failed to parse as int; %s; %s", ds, err)
		}
		sqs = append(sqs, uint(n))
	}
	return sqs
}
