package kata

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func PrimeFactors(n int) string {
	ps := Primes(n)
	return FormatPrimes(ps)
}

func Primes(n int) []int {
	if n == 2 {
		return []int{2}
	}
	p := Find(n)
	if p == n {
		return []int{p}
	} else {
		next := Primes(n / p)
		result := []int{p}
		result = append(result, next...)
		return result
	}
}

func Find(n int) int {
	upper := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < upper; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func FormatPrimes(ps []int) string {
	counter := map[int]int{}
	for _, p := range ps {
		count, ok := counter[p]
		if ok {
			counter[p] = count + 1
		} else {
			counter[p] = 1
		}
	}
	primes := []int{}
	for p := range counter {
		primes = append(primes, p)
	}
	sort.Ints(primes)
	ss := []string{}
	for _, p := range primes {
		count := counter[p]
		var s string
		if count == 1 {
			s = fmt.Sprintf("(%d)", p)
		} else {
			s = fmt.Sprintf("(%d**%d)", p, count)
		}
		ss = append(ss, s)
	}
	return strings.Join(ss, "")
}
