package kata

import "strconv"

func PrevMultOfThree(n int) interface{} {
	orig := strconv.Itoa(n)
	for i := 0; i < len(orig); i++ {
		ds := orig[:len(orig)-i]
		m, _ := strconv.Atoi(ds)
		if m%3 == 0 {
			return m
		}
	}
	return nil
}
