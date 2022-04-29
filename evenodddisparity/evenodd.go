package evenodddisparity

import "strconv"

func EvenOddDisparity(ss []string) int {
	even := 0
	odd := 0
	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err == nil {
			if n%2 == 0 {
				even += 1
			} else {
				odd += 1
			}
		}
	}
	return even - odd
}
