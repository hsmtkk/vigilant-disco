package pascaltriangle

func PascalTriangle(n uint) []uint {
	var m uint = 1
	result := []uint{}
	for m = 1; m <= n; m++ {
		result = append(result, array(m)...)
	}
	return result
}

func array(n uint) []uint {
	if n == 1 {
		return []uint{1}
	}
	result := []uint{1}
	prev := array(n - 1)
	var i uint = 0
	for i = 0; i < n-2; i++ {
		result = append(result, prev[i]+prev[i+1])
	}
	result = append(result, 1)
	return result
}
