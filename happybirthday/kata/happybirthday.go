package kata

func WrapPresent(height, width, length int) int {
	if height <= width && height <= length {
		return 20 + 2*(height+width) + 2*(height+length)
	} else if width <= height && width <= length {
		return 20 + 2*(width+height) + 2*(width+length)
	} else {
		return 20 + 2*(length+height) + 2*(length+width)
	}
}
