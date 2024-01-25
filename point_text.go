package sprint

import "fmt"

func PointText(p Point) Point {
	formattedText := fmt.Sprintf("%s at (%f, %f)", p.Text, p.X, p.Y)
	return MakePoint(p.X, p.Y, formattedText)
}
