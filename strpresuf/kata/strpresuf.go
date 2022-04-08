package kata

func Solve(s string) int {
	for i := len(s) / 2; i > 0; i-- {
		prefix := s[:i]
		suffix := s[len(s)-i:]
		if prefix == suffix {
			return i
		}
	}
	return 0
}
