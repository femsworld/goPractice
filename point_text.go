package sprint

import "fmt"

func PointText(p Point) Point {
	formattedText := fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return Point{X: p.X, Y: p.Y, Text: formattedText}
}
