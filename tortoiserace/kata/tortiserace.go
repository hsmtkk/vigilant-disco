package kata

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	sec := 3600 * g / (v2 - v1)
	min := sec / 60
	sec %= 60
	hour := min / 60
	min %= 60
	return [3]int{hour, min, sec}
}
