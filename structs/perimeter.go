package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}
