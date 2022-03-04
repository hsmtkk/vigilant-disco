package diffvolcube

func CalculateDifference(a, b []uint) uint {
	a3 := Cube(a)
	b3 := Cube(b)
	return Diff(a3, b3)
}

func Cube(a []uint) uint {
	return a[0] * a[1] * a[2]
}

func Diff(a, b uint) uint {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
