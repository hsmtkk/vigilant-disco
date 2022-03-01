package validpin

import "unicode"

func Validate(pin string) bool {
	if len(pin) != 4 && len(pin) != 6 {
		return false
	}
	for _, ch := range pin {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
