package hellonameworld

import "strings"

func Hello(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	head := strings.ToUpper(string(name[0]))
	tail := strings.ToLower(name[1:])
	return "Hello, " + head + tail + "!"
}
