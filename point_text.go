package sprint

import "fmt"

// type Point struct {
// 	X     float32
// 	Y     float32
// 	Text  string
// }

// // MakePoint creates a new Point with the provided values.
// func MakePoint(x, y float32, text string) Point {
// 	return Point{
// 		X:    x,
// 		Y:    y,
// 		Text: text,
// 	}
// }

// PointText formats the text of a Point structure as "Text at (X, Y)".
func PointText(p Point) Point {
	// formattedText := fmt.Sprintf("%s at (%.2f, %.2f)", p.Text, p.X, p.Y)
	formattedText := fmt.Sprintf("%s at (%.2f, %.2f)", p.Text, p.X, p.Y)
	return MakePoint(p.X, p.Y, formattedText)
}
