package kata

func Solve(arr []int) int {
	for _, n := range arr {
		if !contains(arr, -n) {
			return n
		}
	}
	return 0
}

func contains(ns []int, n int) bool {
	for _, m := range ns {
		if m == n {
			return true
		}
	}
	return false
}
