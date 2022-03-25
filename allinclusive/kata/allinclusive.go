package kata

func ContainAllRots(strng string, arr []string) bool {
	for _, rot := range AllRotations(strng) {
		if !Included(arr, rot) {
			return false
		}
	}
	return true
}

func AllRotations(s string) []string {
	results := []string{}
	for i := 0; i < len(s); i++ {
		rotated := s[i:] + s[:i]
		results = append(results, rotated)
	}
	return results
}

func Included(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
