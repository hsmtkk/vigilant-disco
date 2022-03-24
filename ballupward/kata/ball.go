package kata

func MaxBall(v0 int) int {
	v := float64(v0) * 1000 / 3600
	max := 0.0
	maxi := 0
	for i := 0; ; i++ {
		t := 0.1 * float64(i)
		h := calcHeight(v, t)
		if h < 0 {
			break
		}
		if h > max {
			max = h
			maxi = i
		}
	}
	return maxi
}

const GRAVITY = 9.81

func calcHeight(v0, t float64) float64 {
	return v0*t - 0.5*GRAVITY*t*t
}
