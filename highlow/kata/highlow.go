package kata

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	nums := splitToInts(in)
	high, low := findHighLow(nums)
	return fmt.Sprintf("%d %d", high, low)
}

func splitToInts(in string) []int {
	digits := strings.Split(in, " ")
	nums := []int{}
	for _, d := range digits {
		n, err := strconv.Atoi(d)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func findHighLow(nums []int) (int, int) {
	high := nums[0]
	low := nums[0]
	for _, n := range nums {
		if n > high {
			high = n
		}
		if n < low {
			low = n
		}
	}
	return high, low
}
