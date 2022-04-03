package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	x := signature[0]
	y := signature[1]
	z := signature[2]
	results := []float64{x, y, z}
	for i := 0; i < n; i++ {
		newx := y
		newy := z
		newz := x + y + z
		results = append(results, newz)
		x = newx
		y = newy
		z = newz
	}
	results = results[:n]
	return results
}
