package main

import "fmt"

type Shape interface {
	draw()
}

type Rectangle struct {
}

func (r *Rectangle) draw() {
	fmt.Println("Shape: Rectangle")
}

type Circle struct {
}

func (r *Circle) draw() {
	fmt.Println("Shape: Circle")
}

type ShapeDecorator struct {
	decoratedShape Shape
}

func (s *ShapeDecorator) draw() {
	s.decoratedShape.draw()
}

type RedShapeDecorator struct {
	ShapeDecorator
}

func NewRedShapeDecorator(decoratedShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{
		ShapeDecorator{decoratedShape: decoratedShape},
	}
}

func (r *RedShapeDecorator) draw() {
	r.decoratedShape.draw()
	r.setRedBorder()
}

func (r *RedShapeDecorator) setRedBorder() {
	fmt.Println("Border Color: Red")
}
