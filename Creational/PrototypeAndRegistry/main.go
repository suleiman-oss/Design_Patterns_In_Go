package main

import (
	Shapes "github.com/suleiman-oss/Design_Patterns_in_Go/Creational/PrototypeAndRegistry/Shape"
)

func main() {
	circle1 := Shapes.Circle{Radius: 3.0}
	square1 := Shapes.Square{Side: 2}
	registry := Shapes.NewShapeRegistry()
	registry.Register("circle", &circle1)
	registry.Register("square", &square1)
	circle2, err := registry.GetShape("circle")
	if err == nil {
		circle2.GetInfo()
	}
}
