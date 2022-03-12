package arrayleader

func ArrayLeader(array []int) []int {
	result := []int{}
	for i := 0; i < len(array)-1; i++ {
		if array[i] > rightSum(array, i) {
			result = append(result, array[i])
		}
	}
	last := array[len(array)-1]
	if last > 0 {
		result = append(result, last)
	}
	return result
}

func rightSum(array []int, index int) int {
	sum := 0
	for _, n := range array[index+1:] {
		sum += n
	}
	return sum
}
