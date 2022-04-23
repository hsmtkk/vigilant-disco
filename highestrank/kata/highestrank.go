package kata

func HighestRank(nums []int) int {
	counter := map[int]int{}
	for _, n := range nums {
		count, ok := counter[n]
		if ok {
			counter[n] = count + 1
		} else {
			counter[n] = 1
		}
	}
	maxCount := 0
	maxNumber := nums[0]
	for number, count := range counter {
		if count > maxCount {
			maxCount = count
			maxNumber = number
		} else if count == maxCount {
			if number > maxNumber {
				maxNumber = number
			}
		}
	}
	return maxNumber
}
