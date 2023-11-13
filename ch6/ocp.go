package ch6

import "math"

type Shape struct {
	Type   string
	Width  float64
	Height float64
	Radius float64
}

func NewShape(t string) *Shape {
	return &Shape{Type: t}
}

func (s Shape) Area() float64 {
	if s.Type == "rectangle" {
		return s.RectangleArea()
	} else if s.Type == "circle" {
		return s.CircleArea()
	}
	return 0
}

func (s Shape) RectangleArea() float64 {
	return s.Width * s.Height
}

func (s Shape) CircleArea() float64 {
	return math.Pi * math.Pow(s.Radius, 2)
}
