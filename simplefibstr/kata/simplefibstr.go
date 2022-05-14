package kata

func Solve(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "01"
	} else {
		return Solve(n-1) + Solve(n-2)
	}
}
