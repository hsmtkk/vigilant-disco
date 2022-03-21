package kata

import "math"

func Movie(card, ticket int, perc float64) int {
	for i := 1; ; i++ {
		if asystem(ticket, i) > int(math.Ceil(bsystem(card, ticket, perc, i))) {
			return i
		}
	}
}

func asystem(ticket, n int) int {
	return n * ticket
}

func bsystem(card, ticket int, perc float64, n int) float64 {
	return float64(card) + float64(ticket)*perc*(1-math.Pow(perc, float64(n)))/(1-perc)
}
