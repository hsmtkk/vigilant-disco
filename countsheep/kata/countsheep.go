package kata

import (
	"fmt"
	"strings"
)

func countSheep(num int) string {
	ss := []string{}
	for i := 1; i <= num; i++ {
		s := fmt.Sprintf("%d sheep...", i)
		ss = append(ss, s)
	}
	return strings.Join(ss, "")
}
