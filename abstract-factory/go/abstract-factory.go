package main

import "fmt"

type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

type AbstractFactory interface {
	GetShape(shapeType string) Shape
	GetColor(color string) Color
}

// Shape

type Rectangle struct {
}

func (*Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method")
}

type Square struct {
}

func (*Square) Draw() {
	fmt.Println("Inside Square::draw() method")
}

type Circle struct {
}

func (*Circle) Draw() {
	fmt.Println("Inside Circle::draw() method")
}

// --------

// Color

type Red struct {
}

func (*Red) Fill() {
	fmt.Println("Inside Red::fill() method")
}

type Green struct {
}

func (*Green) Fill() {
	fmt.Println("Inside Green::fill() method")
}

type Blue struct {
}

func (*Blue) Fill() {
	fmt.Println("Inside Blue::fill() method")
}

// --------

// ShapeFactory 工厂类
type ShapeFactory struct {
}

func (*ShapeFactory) GetShape(shapeType string) Shape {
	switch shapeType {
	case "RECTANGLE":
		return &Rectangle{}
	case "CIRCLE":
		return &Circle{}
	case "SQUARE":
		return &Square{}
	default:
		return nil
	}
}
func (*ShapeFactory) GetColor(_ string) Color {
	return nil
}

// ColorFactory 工厂类
type ColorFactory struct {
}

func (*ColorFactory) GetShape(_ string) Shape {
	return nil
}

func (*ColorFactory) GetColor(color string) Color {
	switch color {
	case "RED":
		return &Red{}
	case "GREEN":
		return &Green{}
	case "BLUE":
		return &Blue{}
	default:
		return nil
	}
}

// -------

// FactoryProducer 工厂生成器
type FactoryProducer struct {
}

func (*FactoryProducer) GetFactory(choice string) AbstractFactory {
	switch choice {
	case "SHAPE":
		return &ShapeFactory{}
	case "COLOR":
		return &ColorFactory{}
	default:
		return nil
	}
}

// ------
