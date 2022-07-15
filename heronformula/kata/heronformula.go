package kata

import "math"

func Heron(a, b, c int) float32 {
	var s float64 = float64(a+b+c) / 2.0
	area := math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
	return float32(area)
}
