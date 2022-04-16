package kata

import "math"

func Process(n int) int {
	sq := int(math.Sqrt(float64(n)))
	if int(sq*sq) == n {
		return int(sq)
	} else {
		return n * n
	}
}

func SquareOrSquareRoot(arr []int) []int {
	result := []int{}
	for _, n := range arr {
		result = append(result, Process(n))
	}
	return result
}
