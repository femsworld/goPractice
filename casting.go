package sprint

import "math"

func Casting(n float64) int {
	roundedValue := int(math.Round(n))
    return roundedValue
}