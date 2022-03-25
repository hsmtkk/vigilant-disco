package spoonerizeme

import (
	"fmt"
	"strings"
)

func Spoonerize(s string) string {
	words := strings.Split(s, " ")
	first := words[0]
	second := words[1]
	return fmt.Sprintf("%s%s %s%s", string(second[0]), first[1:], string(first[0]), second[1:])
}
