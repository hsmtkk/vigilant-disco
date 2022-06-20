package pairofint

type Pair struct {
	First  int
	Second int
}

func GeneratePairs(m, n int) []Pair {
	result := []Pair{}
	for i := m; i <= n; i++ {
		for j := i; j <= n; j++ {
			result = append(result, Pair{i, j})
		}
	}
	return result
}
