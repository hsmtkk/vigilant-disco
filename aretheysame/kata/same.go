package kata

func Comp(array1 []int, array2 []int) bool {
	squared := map[int]bool{}
	for _, n := range array1 {
		squared[n*n] = true
	}
	for _, n := range array2 {
		_, ok := squared[n]
		if !ok {
			return false
		}
	}
	return true
}
