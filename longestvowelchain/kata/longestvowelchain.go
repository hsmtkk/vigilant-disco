package kata

func Solve(s string) int {
	for i := len(s); i > 0; i-- {
		chain, found := FindVowelChain(s, i)
		if found {
			return len(chain)
		}
	}
	return 0
}

func FindVowelChain(s string, length int) (string, bool) {
	for i := 0; i < len(s)-length; i++ {
		subStr := s[i : i+length]
		if IsVowelChain(subStr) {
			return subStr, true
		}
	}
	return "", false
}

func IsVowelChain(s string) bool {
	for _, ch := range s {
		if ch != 'a' && ch != 'i' && ch != 'u' && ch != 'e' && ch != 'o' {
			return false
		}
	}
	return true
}
