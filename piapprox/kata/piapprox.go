package kata

import (
	"fmt"
	"math"
)

func IterPi(epsilon float64) (int, string) {
	k := 0
	sum := float64(0)
	for {
		sum += calculate(k)
		approxPi := 4 * sum
		if math.Abs(approxPi-math.Pi) < epsilon {
			return k + 1, fmt.Sprintf("%.10f", approxPi)
		}
		k += 1
	}
}

func calculate(k int) float64 {
	if k%2 == 0 {
		return 1 / float64(2*k+1)
	} else {
		return -1 / float64(2*k+1)
	}
}
