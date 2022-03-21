package kata

func CountPositivesSumNegatives(numbers []int) []int {
	count := 0
	sum := 0
	for _, n := range numbers {
		if n > 0 {
			count += 1
		} else if n < 0 {
			sum += n
		}
	}
	return []int{count, sum}
}
