package lifepathnum

import (
	"strconv"
	"strings"
)

func LifePathNumber(birthDate string) int {
	splitted := strings.Split(birthDate, "-")
	year, _ := strconv.Atoi(splitted[0])
	month, _ := strconv.Atoi(splitted[1])
	day, _ := strconv.Atoi(splitted[2])
	return compressNumber(compressNumber(year) + compressNumber(month) + compressNumber(day))
}

func compressNumber(number int) int {
	if number < 10 {
		return number
	}
	s := 0
	for number > 0 {
		s += number % 10
		number /= 10
	}
	return compressNumber(s)
}
