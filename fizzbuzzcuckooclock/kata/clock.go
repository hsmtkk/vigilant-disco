package kata

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	hour, minute, err := parseTime(time)
	if err != nil {
		log.Print(err)
		return ""
	}
	return formatResult(hour, minute)
}

func parseTime(time string) (int, int, error) {
	elems := strings.Split(time, ":")
	if len(elems) != 2 {
		return 0, 0, fmt.Errorf("invalid format; %s", time)
	}
	hourStr := elems[0]
	minuteStr := elems[1]
	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse as int; %s; %s", err, hourStr)
	}
	minute, err := strconv.Atoi(minuteStr)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse as int; %s; %s", err, minuteStr)
	}
	return hour, minute, nil
}

func formatResult(hour, minute int) string {
	if minute == 0 {
		ss := []string{}
		for i := 0; i < hourCount(hour); i++ {
			ss = append(ss, "Cuckoo")
		}
		return strings.Join(ss, " ")
	} else if minute == 30 {
		return "Cuckoo"
	} else if minute%15 == 0 {
		return "Fizz Buzz"
	} else if minute%5 == 0 {
		return "Buzz"
	} else if minute%3 == 0 {
		return "Fizz"
	} else {
		return "tick"
	}
}

func hourCount(hour int) int {
	if hour == 0 {
		return 12
	} else if hour > 12 {
		return hour - 12
	}
	return hour
}
