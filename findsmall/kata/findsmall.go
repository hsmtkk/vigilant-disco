package kata

import "strconv"

func Smallest(n int64) []int64 {
	min := n
	for _, n := range ListCandidates(n) {
		if n < min {

		}
	}
	return nil
}

func ListCandidates(n int64) []int64 {
	results := []int64{}
	ds := strconv.FormatInt(n, 10)
	for i := 0; i < len(ds); i++ {
		rest := ds[:i] + ds[i+1:]
		for j := 0; j < len(rest); j++ {
			s := rest[:j] + string(ds[i]) + rest[j:]
			n, _ := strconv.ParseInt(s, 10, 64)
			results = append(results, n)
		}
		s := rest + string(ds[i])
		n, _ := strconv.ParseInt(s, 10, 64)
		results = append(results, n)
	}
	return results
}
