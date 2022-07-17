package kata

func ReduceFraction(fraction [2]int) [2]int {
	g := gcd(fraction[0], fraction[1])
	return [2]int{fraction[0] / g, fraction[1] / g}
}

func gcd(m, n int) int {
	if m < n {
		return gcd(n, m)
	}
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}
