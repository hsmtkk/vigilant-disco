package secminhour

import "fmt"

func ToTime(seconds uint) string {
	hour := seconds / 3600
	min := (seconds - 3600*hour) / 60
	return fmt.Sprintf("%d hour(s) and %d minute(s)", hour, min)
}
