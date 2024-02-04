package main

import "fmt"

type Shape interface {
	draw() string
}

type Rectangle struct {
}

func (r Rectangle) draw() string {
	fmt.Println("Rectangle")
	return "Rectangle"
}

type Square struct {
}

func (s Square) draw() string {
	fmt.Println("Square")
	return "Square"
}

type ShapeFactory struct {
}

func (s ShapeFactory) Create(shapeType string) Shape {
	var shape Shape
	switch shapeType {
	case "Rectangle":
		shape = &Rectangle{}
		break
	case "Square":
		shape = &Square{}
		break
	default:
		shape = nil
	}
	return shape
}
