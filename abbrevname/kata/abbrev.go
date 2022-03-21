package kata

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	firstLast := strings.Split(name, " ")
	first := firstLast[0]
	last := firstLast[1]
	return fmt.Sprintf("%s.%s", strings.ToUpper(string(first[0])), strings.ToUpper(string(last[0])))
}
