package maxminarray

import "sort"

func MaxMinArray(ns []int) []int {
	half := len(ns) / 2
	sort.Ints(ns)
	mins := ns[:half]
	maxs := ns[half:]
	maxs = reverse(maxs)
	result := []int{}
	for i := 0; i < half; i++ {
		result = append(result, maxs[i])
		result = append(result, mins[i])
	}
	if len(ns)%2 == 1 {
		result = append(result, maxs[half])
	}
	return result
}

func reverse(ns []int) []int {
	reversed := []int{}
	for _, n := range ns {
		reversed = append([]int{n}, reversed...)
	}
	return reversed
}
