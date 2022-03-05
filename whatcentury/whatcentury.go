package whatcentury

import "fmt"

func WhatCentury(year uint32) string {
	century := (year + 99) / 100
	return fmt.Sprintf("%d%s", century, suffix(century))
}

func suffix(century uint32) string {
	if century < 20 {
		return "th"
	}
	switch century % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}
