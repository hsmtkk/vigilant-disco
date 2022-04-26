package kata

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	height := 0
	for i := 0; ; i++ {
		height += upSpeed
		if height >= desiredHeight {
			return i + 1
		}
		height -= downSpeed
	}
}
