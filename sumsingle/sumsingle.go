package sumsingle

func SumSingle(ns []int) int {
	counter := map[int]int{}
	for _, n := range ns {
		count, ok := counter[n]
		if ok {
			counter[n] = count + 1
		} else {
			counter[n] = 1
		}
	}
	sum := 0
	for n, count := range counter {
		if count == 1 {
			sum += n
		}
	}
	return sum
}
