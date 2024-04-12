package Shapes

import "errors"

type ShapeRegistry struct {
	shapes map[string]Shape
}

func NewShapeRegistry() *ShapeRegistry {
	return &ShapeRegistry{shapes: map[string]Shape{}}
}

func (sr *ShapeRegistry) Register(name string, shape Shape) {
	sr.shapes[name] = shape
}

func (sr *ShapeRegistry) GetShape(name string) (Shape, error) {
	if sp, ok := sr.shapes[name]; ok {
		return sp.GetClone(), nil
	}
	return nil, errors.New("Shape with name not registered")
}
