package simplearrayproduct

import "sort"

func Solve(arrays ...[]int) int {
	return max(solve2(arrays))
}

func max(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func solve2(arrays [][]int) []int {
	if len(arrays) == 1 {
		return arrays[0]
	}
	next := solve2(arrays[1:])
	result := []int{}
	for _, x := range arrays[0] {
		for _, y := range next {
			result = append(result, x*y)
		}
	}
	return result
}
