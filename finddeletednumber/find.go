package finddeletednumber

func Find(list []int, mixedList []int) (bool, int) {
	nums := map[int]bool{}
	for _, n := range list {
		nums[n] = true
	}
	for _, n := range mixedList {
		delete(nums, n)
	}
	for n := range nums {
		return true, n
	}
	return false, 0
}
