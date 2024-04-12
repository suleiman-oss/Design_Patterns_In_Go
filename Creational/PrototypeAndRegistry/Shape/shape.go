package Shapes

import "fmt"

type Shape interface {
	GetClone() Shape
	GetInfo()
}

type Circle struct {
	Radius float32
}

func (c *Circle) GetClone() Shape {
	return &Circle{Radius: c.Radius}
}

func (c *Circle) GetInfo() {
	fmt.Printf("Circle with radius %f\n", c.Radius)
}

type Square struct {
	Side float32
}

func (s *Square) GetClone() Shape {
	return &Square{Side: s.Side}
}

func (s *Square) GetInfo() {
	fmt.Printf("Square with side %f\n", s.Side)
}
