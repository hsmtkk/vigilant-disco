package kata

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	a1Max := maxLen(a1)
	a1Min := minLen(a1)
	a2Max := maxLen(a2)
	a2Min := minLen(a2)
	abs1 := abs(a1Max - a2Min)
	abs2 := abs(a2Max - a1Min)
	if abs1 > abs2 {
		return abs1
	} else {
		return abs2
	}
}

func maxLen(ss []string) int {
	max := 0
	for _, s := range ss {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}

func minLen(ss []string) int {
	min := len(ss[0])
	for _, s := range ss {
		if len(s) < min {
			min = len(s)
		}
	}
	return min
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}
