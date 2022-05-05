package boyman

import "sort"

func MenFromBoys(ns []int) []int {
	evenUniq := map[int]bool{}
	oddUniq := map[int]bool{}
	for _, n := range ns {
		if n%2 == 0 {
			evenUniq[n] = true
		} else {
			oddUniq[n] = true
		}
	}
	even := []int{}
	odd := []int{}
	for n := range evenUniq {
		even = append(even, n)
	}
	for n := range oddUniq {
		odd = append(odd, n)
	}
	sort.Ints(even)
	sort.Ints(odd)
	odd = reverse(odd)
	even = append(even, odd...)
	return even
}

func reverse(ns []int) []int {
	reversed := []int{}
	for _, n := range ns {
		reversed = append([]int{n}, reversed...)
	}
	return reversed
}
