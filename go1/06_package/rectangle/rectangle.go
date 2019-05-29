package rectangle

import (
	"math"
)

// all methods with a capital letter at the beginning are exported functions
func Area(length, width float64) float64 {
	return length * width
}

func Diagonal(length, width float64) float64 {
	return math.Sqrt(math.Pow(length, 2) + math.Pow(width, 2))
}
