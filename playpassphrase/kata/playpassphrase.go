package kata

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func PlayPass(s string, n int) string {
	return Step5(Step4(Step2(Step1(s, n))))
}

func Step1(s string, n int) string {
	result := []string{}
	for _, ch := range s {
		if 'a' <= ch && ch <= 'z' {
			shifted := ch + rune(n)
			if shifted > 'z' {
				shifted -= 26
			}
			result = append(result, string(shifted))
		} else if 'A' <= ch && ch <= 'Z' {
			shifted := ch + rune(n)
			if shifted > 'Z' {
				shifted -= 26
			}
			result = append(result, string(shifted))
		} else {
			result = append(result, string(ch))
		}
	}
	return strings.Join(result, "")
}

func Step2(s string) string {
	result := []string{}
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			d, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatal(err)
			}
			n := 9 - d
			result = append(result, strconv.Itoa(n))
		} else {
			result = append(result, string(ch))
		}
	}
	return strings.Join(result, "")
}

func Step4(s string) string {
	result := []string{}
	for i := 0; i < len(s); i++ {
		converted := string(s[i])
		if i%2 == 0 {
			if isLower(converted) {
				converted = strings.ToUpper(converted)
			}
		} else {
			if isUpper(converted) {
				converted = strings.ToLower(converted)
			}
		}
		result = append(result, converted)
	}
	return strings.Join(result, "")
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func Step5(s string) string {
	reversed := []string{}
	for _, r := range s {
		reversed = append([]string{string(r)}, reversed...)
	}
	return strings.Join(reversed, "")
}
