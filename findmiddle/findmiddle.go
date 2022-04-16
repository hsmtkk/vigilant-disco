package findmiddle

import (
	"sort"
)

func FindMiddle(ns []int) int {
	sorted := append([]int{}, ns...)
	sort.Ints(sorted)
	middle := sorted[1]
	for i, n := range ns {
		if n == middle {
			return i
		}
	}
	return 0
}
