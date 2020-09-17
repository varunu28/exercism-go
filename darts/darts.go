package darts

import (
	"math"
)

const (
	inner  = 1.0
	middle = 5.0
	outer  = 10.0
)

// Score returns score of a dart attempt
func Score(x float64, y float64) int {
	hypot := math.Hypot(x, y)
	switch {
	case hypot <= inner:
		return 10
	case hypot <= middle:
		return 5
	case hypot <= outer:
		return 1
	default:
		return 0
	}
}
