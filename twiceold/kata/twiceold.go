package kata

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	return int(math.Abs(float64(2*sonYearsOld - dadYearsOld)))
}
