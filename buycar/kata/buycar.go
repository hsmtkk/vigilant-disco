package kata

import "math"

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	if startPriceOld >= startPriceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}
	old := float64(startPriceOld)
	new := float64(startPriceNew)
	lossPercent := percentLossByMonth
	for month := 1; ; month++ {
		if month%2 == 0 {
			lossPercent = percentLossByMonth + 0.5*float64(month/2)
		}
		old *= 1 - 0.01*lossPercent
		new *= 1 - 0.01*lossPercent
		saving := savingperMonth * month
		diff := old + float64(saving) - new
		if diff > 0 {
			return [2]int{month, int(math.Round(diff))}
		}
	}
}
