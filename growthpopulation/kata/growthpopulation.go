package kata

func NbYear(p0 int, percent float64, aug int, p int) int {
	current := p0
	year := 0
	for {
		year += 1
		current += int(float64(current)*percent*0.01) + aug
		if current >= p {
			break
		}
	}
	return year
}
