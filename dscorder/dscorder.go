package dscorder

import (
	"sort"
	"strconv"
	"strings"
)

func DscOrder(input uint) uint {
	s := strconv.Itoa(int(input))
	ss := stringArray(s)
	sort.Strings(ss)
	ss = reverse(ss)
	s = strings.Join(ss, "")
	n, _ := strconv.Atoi(s)
	return uint(n)
}

func stringArray(s string) []string {
	result := []string{}
	for _, r := range s {
		result = append(result, string(r))
	}
	return result
}

func reverse(orig []string) []string {
	reversed := []string{}
	for _, s := range orig {
		reversed = append([]string{s}, reversed...)
	}
	return reversed
}
