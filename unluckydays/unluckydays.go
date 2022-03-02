package unluckydays

import "time"

type Calculator interface {
	Calculate(year uint) uint
}

func New() Calculator {
	return &calculatorImpl{}
}

type calculatorImpl struct{}

func (c *calculatorImpl) Calculate(year uint) uint {
	var count uint = 0
	for month := 1; month <= 12; month++ {
		m := time.Month(month)
		d := time.Date(int(year), m, 13, 0, 0, 0, 0, time.UTC)
		if d.Weekday() == time.Friday {
			count += 1
		}
	}
	return count
}
