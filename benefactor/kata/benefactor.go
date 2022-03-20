package kata

import "math"

func NewAvg(arr []float64, navg float64) int64 {
	n := float64(len(arr) + 1)
	ans := n*navg - sum(arr)
	if ans < 0 {
		return -1
	}
	return int64(math.Ceil(ans))
}

func sum(arr []float64) float64 {
	var s float64 = 0
	for _, a := range arr {
		s += a
	}
	return s
}
