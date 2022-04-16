package kata

func Gps(s int, x []float64) int {
	speeds := []float64{}
	for i := 1; i < len(x); i++ {
		delta := x[i] - x[i-1]
		v := (3600 * delta) / float64(s)
		speeds = append(speeds, v)
	}
	return max(speeds)
}

func max(fs []float64) int {
	m := int(fs[0])
	for _, f := range fs {
		if int(f) > m {
			m = int(f)
		}
	}
	return m
}
