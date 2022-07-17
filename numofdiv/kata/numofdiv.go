package kata

func Divisions(n, divisor int) int {
	for i := 0; ; i++ {
		if n < divisor {
			return i
		}
		n /= divisor
	}
}
