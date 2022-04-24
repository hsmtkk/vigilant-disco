package kata

import "strconv"

func LateRide(n int) int {
	hour := n / 60
	minute := n % 60
	return sumAsDigit(hour) + sumAsDigit(minute)
}

func sumAsDigit(n int) int {
	sum := 0
	for _, d := range strconv.Itoa(n) {
		m, _ := strconv.Atoi(string(d))
		sum += m
	}
	return sum
}
