package countip

import (
	"log"
	"strconv"
	"strings"
)

func New() *IPCounter {
	return &IPCounter{}
}

type IPCounter struct{}

func (cnt *IPCounter) CountIP(from, to string) int {
	begin := IP2Int(from)
	end := IP2Int(to)
	return end - begin
}

func IP2Int(ip string) int {
	elems := strings.Split(ip, ".")
	nums := []int{}
	for _, elem := range elems {
		n, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	return nums[0]<<24 + nums[1]<<16 + nums[2]<<8 + nums[3]
}
