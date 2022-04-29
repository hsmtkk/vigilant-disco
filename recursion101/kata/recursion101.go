package kata

func Solve(a, b int) []int {
	/*
			i) If a = 0 or b = 0, return [a,b]. Otherwise, go to step (ii);
		ii) If a ≥ 2*b, set a = a - 2*b, and repeat step (i). Otherwise, go to step (iii);
		iii) If b ≥ 2*a, set b = b - 2*a, and repeat step (i). Otherwise, return [a,b].
	*/
	if a == 0 || b == 0 {
		return []int{a, b}
	} else if a >= 2*b {
		Solve(a-2*b, b)
	} else if b >= 2*a {
		Solve(a, b-2*a)
	}
	return []int{a, b}
}
