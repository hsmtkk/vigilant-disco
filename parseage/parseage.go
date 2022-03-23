package parseage

import (
	"log"
	"strconv"
	"strings"
)

func ParseAge(s string) uint {
	elems := strings.Split(s, " ")
	n, err := strconv.ParseUint(elems[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return uint(n)
}
