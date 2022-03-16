package mostfrequentdays

import (
	"time"
)

func MostFrequentDays(year int) []string {
	d := time.Date(year, time.January, 1, 1, 0, 0, 0, time.UTC)
	end := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	weekDayCounter := map[time.Weekday]uint{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}
	for {
		weekDayCounter[d.Weekday()] += 1
		d = d.AddDate(0, 0, 1)
		if d.After(end) {
			break
		}
	}
	var max uint = 0
	for _, count := range weekDayCounter {
		if count > max {
			max = count
		}
	}
	result := []string{}
	for _, wd := range []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday} {
		count := weekDayCounter[wd]
		if count == max {
			result = append(result, wd.String())
		}
	}
	return result
}
