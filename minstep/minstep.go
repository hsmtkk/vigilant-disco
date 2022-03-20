package minstep

import (
	"log"
	"sort"
)

func MinimumSteps(ns []int, threshold int) int {
	log.Print(ns, threshold)
	for _, n := range ns {
		if n >= threshold {
			return 0
		}
	}
	small1, small2, rest := stripSmall(ns)
	rest = append(rest, small1+small2)
	return 1 + MinimumSteps(rest, threshold)
}

func stripSmall(ns []int) (int, int, []int) {
	sort.Ints(ns)
	return ns[0], ns[1], ns[2:]
	//return ns[len(ns)-1], ns[len(ns)-2], ns[:len(ns)-2]
}
