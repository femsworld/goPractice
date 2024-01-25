package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	pi := float32(3.14)
	area := pi * r * r
	diameter := 2 * r
	perimeter := 2 * r
	return Circle{
		Radius:    r,
		Diameter:  diameter,
		Area:      area,
		Perimeter: perimeter,
	}

}
