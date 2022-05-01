package kata

func SplitAndAdd(numbers []int, n int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	for i := 0; i < n; i++ {
		if len(result) == 1 {
			break
		}
		left, right := split(result)
		result = add(left, right)
	}
	return result
}

func split(numbers []int) ([]int, []int) {
	length := len(numbers)
	center := length / 2
	return numbers[:center], numbers[center:]
}

func add(left, right []int) []int {
	result := []int{}
	if len(left) == len(right) {
		for i := range left {
			result = append(result, left[i]+right[i])
		}
	} else {
		result = append(result, right[0])
		for i := range left {
			result = append(result, left[i]+right[i+1])
		}
	}
	return result
}
