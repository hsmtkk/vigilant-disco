package kata

func SumEvenFibonacci(limit int) int {
	sum := 0
	for _, x := range fibonacciSequence(limit) {
		if x%2 == 0 {
			sum += x
		}
	}
	return sum
}

func fibonacciSequence(limit int) []int {
	x := 1
	y := 2
	sequence := []int{1, 2}
	for {
		z := x + y
		if z > limit {
			return sequence
		}
		sequence = append(sequence, z)
		x = y
		y = z
	}
}
