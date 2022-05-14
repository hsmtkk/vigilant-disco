package foreigncountry

type Span struct {
	Begin int
	End   int
}

func DaysRepresented(trips []Span) int {
	days := 0
	for d := 1; d <= 365; d++ {
		if inRange(d, trips) {
			days += 1
		}
	}
	return days
}

func inRange(d int, trips []Span) bool {
	for _, span := range trips {
		if span.Begin <= d && d <= span.End {
			return true
		}
	}
	return false
}
