package opseq

type Calculator interface {
	Calculate([]int) int
}

func New() Calculator {
	return &calculatorImpl{}
}

type calculatorImpl struct{}

func (c *calculatorImpl) Calculate(ns []int) int {
	sum := 0
	for i, x := range ns {
		if x > 0 {
			x = x * x
		}
		if (i+1)%3 == 0 {
			x = 3 * x
		}
		if (i+1)%5 == 0 {
			x = -x
		}
		sum += x
	}
	return sum
}
