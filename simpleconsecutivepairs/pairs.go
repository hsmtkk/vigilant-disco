package simpleconsecutivepairs

func Count(ns []int) int {
	pairs := splitPairs(ns)
	pairs = selectConsecutive(pairs)
	return len(pairs)
}

type pair struct {
	first  int
	second int
}

func splitPairs(ns []int) []pair {
	pairs := []pair{}
	for i := 0; i < len(ns)-1; i += 2 {
		pairs = append(pairs, pair{ns[i], ns[i+1]})
	}
	return pairs
}

func selectConsecutive(origs []pair) []pair {
	result := []pair{}
	for _, p := range origs {
		x := p.first
		y := p.second
		if x+1 == y || x-1 == y {
			result = append(result, p)
		}
	}
	return result
}
