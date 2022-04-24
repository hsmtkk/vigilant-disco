package kata

func CartesianNeighbor(x, y int) [][]int {
	neighbors := [][]int{}
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			neighbors = append(neighbors, []int{x + dx, y + dy})
		}
	}
	return neighbors
}
