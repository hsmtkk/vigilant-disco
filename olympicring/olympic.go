package olympicring

func Score(s string) string {
	ns := convertInts(s)
	su := sum(ns) / 2
	switch su {
	case 0:
		return "Not even a medal!"
	case 1:
		return "Not even a medal!"
	case 2:
		return "Bronze!"
	case 3:
		return "Silver!"
	default:
		return "Gold!"
	}
}

func convertInts(s string) []int {
	ns := []int{}
	for _, ch := range s {
		ns = append(ns, convertInt(ch))
	}
	return ns
}

func convertInt(ch rune) int {
	switch ch {
	case 'a':
		return 1
	case 'b':
		return 1
	case 'd':
		return 1
	case 'e':
		return 1
	case 'g':
		return 1
	case 'o':
		return 1
	case 'p':
		return 1
	case 'q':
		return 1
	case 'A':
		return 1
	case 'B':
		return 2
	case 'D':
		return 1
	case 'O':
		return 1
	case 'P':
		return 1
	case 'Q':
		return 1
	case 'R':
		return 1
	default:
		return 0
	}
}

func sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}
