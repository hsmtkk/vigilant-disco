package kata

import "math"

const PI = 3.141592653589793 // Use this value as pi in your calculations if necessary

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	return 0.5 * float64(numberOfSides) * circleRadius * circleRadius * math.Sin(PI/float64(numberOfSides))
}
