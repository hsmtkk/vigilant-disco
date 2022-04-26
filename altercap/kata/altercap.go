package kata

import "strings"

func Capitalize(st string) []string {
	return []string{evenCapitalize(st), oddCapitalize(st)}
}

func evenCapitalize(s string) string {
	result := []string{}
	for i := 0; i < len(s); i++ {
		t := string(s[i])
		if i%2 == 0 {
			result = append(result, strings.ToUpper(t))
		} else {
			result = append(result, t)
		}
	}
	return strings.Join(result, "")
}

func oddCapitalize(s string) string {
	result := []string{}
	for i := 0; i < len(s); i++ {
		t := string(s[i])
		if i%2 == 1 {
			result = append(result, strings.ToUpper(t))
		} else {
			result = append(result, t)
		}
	}
	return strings.Join(result, "")
}
