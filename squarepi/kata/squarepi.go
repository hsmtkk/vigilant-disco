package kata

import (
	"log"
	"math"
	"strconv"
)

const PiStr = "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"

func SquarePi(digits int) int {
	sum := 0
	for i := 1; i <= digits; i++ {
		d := PiStr[i-1]
		n, err := strconv.Atoi(string(d))
		if err != nil {
			log.Fatal(err)
		}
		sum += n * n
	}
	return int(math.Ceil(math.Sqrt(float64(sum))))
}
