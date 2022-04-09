package billardstriangle

func Solve(balls int) int {
	for i := 1; ; i++ {
		m := i * (i + 1) / 2
		if m == balls {
			return i
		} else if m > balls {
			return i - 1
		}
	}
}
