package kata

import "strings"

func bandNameGenerator(word string) string {
	head := string(word[0])
	tail := word[1:]
	capitalized := strings.ToUpper(head) + tail
	if strings.EqualFold(head, string(word[len(word)-1])) {
		return capitalized + tail
	} else {
		return "The " + capitalized
	}
}
