package poetpendulum

import "sort"

func Arrange(nums []int) []int {
	sort.Ints(nums)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			result = append([]int{nums[i]}, result...)
		} else {
			result = append(result, nums[i])
		}
	}
	return result
}
