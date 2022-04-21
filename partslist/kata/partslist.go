package kata

import "strings"

func PartList(arr []string) string {
	elems := []string{}
	for i := 1; i < len(arr); i++ {
		left := strings.Join(arr[:i], " ")
		right := strings.Join(arr[i:], " ")
		elems = append(elems, left+", "+right)
	}
	elems = putParen(elems)
	return strings.Join(elems, "")
}

func putParen(ss []string) []string {
	result := []string{}
	for _, s := range ss {
		result = append(result, "("+s+")")
	}
	return result
}
