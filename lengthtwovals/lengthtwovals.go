package lengthtwovals

func Alternate(length int, first, second string) []string {
	result := []string{}
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			result = append(result, first)
		} else {
			result = append(result, second)
		}
	}
	return result
}
