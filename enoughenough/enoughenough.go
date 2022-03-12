package enoughenough

func DeleteNth(numbers []uint, threshold uint) []uint {
	result := []uint{}
	numberCounter := map[uint]uint{}
	for _, n := range numbers {
		count, ok := numberCounter[n]
		if ok {
			numberCounter[n] = count + 1
		} else {
			numberCounter[n] = 1
		}
		count = numberCounter[n]
		if count <= threshold {
			result = append(result, n)
		}
	}
	return result
}
