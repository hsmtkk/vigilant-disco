package correcttimestr

import (
	"fmt"
	"strconv"
	"strings"
)

func Format(orig string) (string, error) {
	splitted := strings.Split(orig, ":")
	if len(splitted) != 3 {
		return "", fmt.Errorf("invalid format")
	}
	hour, min, sec, err := parseInt(splitted)
	if err != nil {
		return "", err
	}
	min += sec / 60
	sec %= 60
	hour += min / 60
	min %= 60
	hour %= 24
	return fmt.Sprintf("%02d:%02d:%02d", hour, min, sec), nil
}

func parseInt(ss []string) (int, int, int, error) {
	h, err := strconv.Atoi(ss[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("not number: %s", ss[0])
	}
	m, err := strconv.Atoi(ss[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("not number: %s", ss[1])
	}
	s, err := strconv.Atoi(ss[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("not number: %s", ss[2])
	}
	return h, m, s, nil
}
