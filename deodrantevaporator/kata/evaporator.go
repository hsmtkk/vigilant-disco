package kata

func Evaporator(content float64, evapPerDay int, threshold int) int {
	thresholdContent := content * 0.01 * float64(threshold)
	days := 0
	for {
		days += 1
		content *= 1 - 0.01*float64(evapPerDay)
		if content < thresholdContent {
			break
		}
	}
	return days
}
