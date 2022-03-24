package multiplicationtable

func MultiplicationTable(size int) [][]int {
	result := [][]int{}
	for row := 0; row < size; row++ {
		line := []int{}
		for column := 0; column < size; column++ {
			line = append(line, (row+1)*(column+1))
		}
		result = append(result, line)
	}
	return result
}
