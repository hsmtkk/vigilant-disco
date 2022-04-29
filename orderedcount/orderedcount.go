package orderedcount

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	tuples := []Tuple{}
	for _, ch := range text {
		found, index := indexOf(ch, tuples)
		if found {
			tuples[index].Count += 1
		} else {
			tuples = append(tuples, Tuple{ch, 1})
		}
	}
	return tuples
}

func indexOf(char rune, tuples []Tuple) (bool, int) {
	for i, t := range tuples {
		if char == t.Char {
			return true, i
		}
	}
	return false, 0
}
