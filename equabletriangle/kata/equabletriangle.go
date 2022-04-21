package kata

import "math"

func EquableTriangle(a, b, c int) bool {
	return area(a, b, c) == perimeter(a, b, c)
}

func area(a, b, c int) float64 {
	s := float64(a+b+c) * 0.5
	return math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
}

func perimeter(a, b, c int) float64 {
	return float64(a + b + c)
}
