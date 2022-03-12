package summin

func SumMin(numbers [][]uint) uint {
	mins := []uint{}
	for _, row := range numbers {
		mins = append(mins, selectMin(row))
	}
	return sum(mins)
}

func selectMin(numbers []uint) uint {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func sum(ns []uint) uint {
	var s uint = 0
	for _, n := range ns {
		s += n
	}
	return s
}
