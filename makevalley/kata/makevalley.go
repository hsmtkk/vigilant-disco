package kata

import "sort"

func MakeValley(arr []int) []int {
	sorted := append([]int{}, arr...)
	sort.Ints(sorted)
	result := []int{sorted[0]}
	for i := 1; i < len(sorted); i++ {
		if i%2 == 0 {
			result = append([]int{sorted[i]}, result...)
		} else {
			result = append(result, sorted[i])
		}
	}
	return result
}
