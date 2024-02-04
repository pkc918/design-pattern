package main

import "testing"

func TestShapeFactory_Create(t *testing.T) {
	shapeFactory := &ShapeFactory{}
	rectangle := shapeFactory.Create("Rectangle")
	square := shapeFactory.Create("Square")

	if rectangle.draw() != "Rectangle" && square.draw() != "Square" {
		t.Fatal("shape factory test fail")
	}
}
