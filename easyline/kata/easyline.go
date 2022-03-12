package kata

import "math/big"

func EasyLine(n int) string {
	bn := EasyLineNum(n)
	return bn.String()
}

func EasyLineNum(n int) *big.Int {
	ns := EasyLineNumbers(n)
	return SumOfSquares(ns)
}

func EasyLineNumbers(n int) []*big.Int {
	if n == 0 {
		return []*big.Int{big.NewInt(1)}
	}
	prev := EasyLineNumbers(n - 1)
	next := []*big.Int{big.NewInt(1)}
	for i := 0; i < n-1; i++ {
		x := prev[i]
		x = x.Add(x, prev[i+1])
		next = append(next, x)
	}
	next = append(next, big.NewInt(1))
	return next
}

func SumOfSquares(ns []*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, n := range ns {
		x := n.Mul(n, n)
		sum = sum.Add(sum, x)
	}
	return sum
}
