package printerror

import "fmt"

func PrintError(s string) string {
	all := len(s)
	error := countError(s)
	return fmt.Sprintf("%d/%d", error, all)
}

func countError(s string) uint {
	var count uint = 0
	for _, ch := range s {
		if 'm' < ch {
			count += 1
		}
	}
	return count
}
