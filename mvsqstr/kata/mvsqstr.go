package kata

import "strings"

func VertMirror(s string) string {
	rows := strings.Split(s, "\n")
	result := []string{}
	for _, row := range rows {
		result = append(result, reverseString(row))
	}
	return strings.Join(result, "\n")
}

func HorMirror(s string) string {
	rows := strings.Split(s, "\n")
	reversed := reverseSlice(rows)
	return strings.Join(reversed, "\n")
}

func reverseString(s string) string {
	result := ""
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

func reverseSlice(ss []string) []string {
	reversed := []string{}
	for _, s := range ss {
		reversed = append([]string{s}, reversed...)
	}
	return reversed
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}
