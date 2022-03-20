package kata

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Gcdi(x, y int) int {
	return gcd(abs(x), abs(y))
}

func gcd(x, y int) int {
	start := x
	if x > y {
		start = y
	}
	for i := start; i > 0; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1
}

func Som(x, y int) int {
	return x + y
}

func Maxi(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Mini(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func Lcmu(x, y int) int {
	return lcm(abs(x), abs(y))
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	results := []int{}
	acc := init
	for _, x := range arr {
		acc = f(acc, x)
		results = append(results, acc)
	}
	return results
}
