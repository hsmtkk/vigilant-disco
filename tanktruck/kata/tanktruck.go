package kata

import "math"

func TankVol(h, d, vt int) int {
	hf := float64(h)
	r := float64(d) * 0.5
	theta := math.Acos((r - hf) / r)
	v := float64(vt) * (r*r*theta - r*(r-hf)*math.Sin(theta)) / (math.Pi * r * r)
	return int(v)
}
