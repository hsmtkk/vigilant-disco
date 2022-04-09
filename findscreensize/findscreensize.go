package findscreensize

import (
	"fmt"
	"strconv"
	"strings"
)

func FindScreenSize(width int, ratio string) string {
	elems := strings.Split(ratio, ":")
	ratioWidth, _ := strconv.Atoi(elems[0])
	ratioHeight, _ := strconv.Atoi(elems[1])
	height := width * ratioHeight / ratioWidth
	return fmt.Sprintf("%dx%d", width, height)
}
