package kata

type Brace int

const (
	Round Brace = iota
	Square
	Curly
)

func ValidBraces(str string) bool {
	braces := []Brace{}
	for _, ch := range str {
		switch ch {
		case '(':
			braces = append(braces, Round)
		case ')':
			if len(braces) == 0 {
				return false
			}
			poped := braces[len(braces)-1]
			braces = braces[:len(braces)-1]
			if poped != Round {
				return false
			}
		case '[':
			braces = append(braces, Square)
		case ']':
			if len(braces) == 0 {
				return false
			}
			poped := braces[len(braces)-1]
			braces = braces[:len(braces)-1]
			if poped != Square {
				return false
			}
		case '{':
			braces = append(braces, Curly)
		case '}':
			if len(braces) == 0 {
				return false
			}
			poped := braces[len(braces)-1]
			braces = braces[:len(braces)-1]
			if poped != Curly {
				return false
			}
		default:
			return false
		}
	}
	return len(braces) == 0
}
