package darts

import "math"

func Score(x, y float64) int {
	dst := math.Sqrt(x*x + y*y)

	switch {
	case dst <= 1:
		return 10
	case dst <= 5:
		return 5
	case dst <= 10:
		return 1
	default:
		return 0
	}
}
