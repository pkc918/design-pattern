package main

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	circle := Circle{}
	redCircle := NewRedShapeDecorator(&Circle{})
	redRectangle := NewRedShapeDecorator(&Rectangle{})
	fmt.Println("Circle with normal border")
	circle.draw()
	fmt.Println("Circle of red border")
	redCircle.draw()
	fmt.Println("Rectangle of red border")
	redRectangle.draw()
}
