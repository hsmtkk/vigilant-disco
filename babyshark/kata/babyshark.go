package kata

import (
	"fmt"
	"strings"
)

const lh = "Let's go hunt"

func BabySharkLyrics() string {
	lines := []string{}
	doo6 := strings.Repeat("doo ", 5) + "doo"
	for _, shark := range []string{"Baby", "Mommy", "Daddy", "Grandma", "Grandpa"} {
		s1 := fmt.Sprintf("%s shark, %s", shark, doo6)
		for i := 0; i < 3; i++ {
			lines = append(lines, s1)
		}
		lines = append(lines, fmt.Sprintf("%s shark!", shark))
	}
	for i := 0; i < 3; i++ {
		lines = append(lines, lh+", "+doo6)
	}
	lines = append(lines, lh+"!")
	lines = append(lines, "Run away,…")
	return strings.Join(lines, "\n") + "\n"
}
