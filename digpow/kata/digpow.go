package kata

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	dp := calculateDigPow(n, p)
	for k := 1; ; k++ {
		var product int64 = int64(n) * int64(k)
		if product == dp {
			return k
		} else if product > dp {
			return -1
		}
	}
}

func calculateDigPow(n, p int) int64 {
	ds := strconv.Itoa(n)
	var s int64 = 0
	for i := 0; i < len(ds); i++ {
		m, _ := strconv.Atoi(string(ds[i]))
		s += int64(math.Pow(float64(m), float64(i+p)))
	}
	return s
}
