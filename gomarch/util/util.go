package util

import "math"

func Saturate(x float64) float64 {
	return math.Max(math.Min(x, 1.0), 0.0)
}
