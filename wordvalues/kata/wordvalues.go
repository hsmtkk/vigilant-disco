package kata

func NameValue(my_list []string) []int {
	values := []int{}
	for i, word := range my_list {
		values = append(values, (i+1)*wordScore(word))
	}
	return values
}

func wordScore(word string) int {
	s := 0
	for _, ch := range word {
		if ch == ' ' {
			continue
		}
		s += charScore(ch)
	}
	return s
}

func charScore(ch rune) int {
	return int(ch-'a') + 1
}
