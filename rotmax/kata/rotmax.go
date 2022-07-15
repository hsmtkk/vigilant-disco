package kata

import (
	"strconv"
)

func MaxRot(n int64) int64 {
	nums := []int64{}
	current := n
	size := len(strconv.FormatInt(n, 10))
	for i := 0; i < size; i++ {
		next := rotate(current, i)
		nums = append(nums, next)
		current = next
	}
	return max(nums)
}

func rotate(n int64, fix int) int64 {
	ds := strconv.FormatInt(n, 10)
	left := ds[:fix]
	right := ds[fix:]
	ds = left + right[1:] + string(right[0])
	m, _ := strconv.ParseInt(ds, 10, 64)
	return m
}

func max(ns []int64) int64 {
	m := ns[0]
	for _, n := range ns {
		if n > m {
			m = n
		}
	}
	return m
}
