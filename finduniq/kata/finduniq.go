package kata

func FindUniq(arr []float32) float32 {
	counter := map[float32]uint{}
	for _, f := range arr {
		count, ok := counter[f]
		if ok {
			counter[f] = count + 1
		} else {
			counter[f] = 1
		}
	}
	for f, count := range counter {
		if count == 1 {
			return f
		}
	}
	return 0
}
