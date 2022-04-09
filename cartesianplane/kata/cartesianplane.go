package kata

type binaryOperator func(int, int) int

func SumFunc(m int, op binaryOperator) int64 {
	var s int64 = 0
	for x := 1; x <= m; x++ {
		for y := 1; y <= m; y++ {
			s += int64(op(x, y))
		}
	}
	return s
}

func SuMin(m int) int64 {
	return SumFunc(m, min)
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func SuMax(m int) int64 {
	return SumFunc(m, max)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func SumSum(m int) int64 {
	return SumFunc(m, add)
}

func add(x, y int) int {
	return x + y
}
