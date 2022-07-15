package kata

func HasUniqueChar(str string) bool {
	chars := map[rune]bool{}
	for _, ch := range str {
		_, ok := chars[ch]
		if ok {
			return false
		}
		chars[ch] = true
	}
	return true
}
