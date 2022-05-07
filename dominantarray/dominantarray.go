package dominantarray

func Solve(ns []int) []int {
	result := []int{}
	for i := 0; i < len(ns)-1; i++ {
		if isGreaterThan(ns[i], ns[i+1:]) {
			result = append(result, ns[i])
		}
	}
	result = append(result, ns[len(ns)-1])
	return result
}

func isGreaterThan(n int, ns []int) bool {
	for _, m := range ns {
		if n <= m {
			return false
		}
	}
	return true
}
