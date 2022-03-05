package kata

import (
	"time"
)

func DateNbDays(a0 int, a int, p int) string {
	dailyInterest := float64(p) / 36000
	days := 0
	current := float64(a0)
	for {
		current = current * (1 + dailyInterest)
		days += 1
		if current >= float64(a) {
			break
		}
	}
	start := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, days)
	return end.Format("2006-01-02")
}
