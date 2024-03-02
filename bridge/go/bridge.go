package main

import "fmt"

type DrawAPI interface {
	drawCircle(radius, x, y int)
}

type RedCircle struct {
}

func (*RedCircle) drawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle [color: red, radius: %v, x: %v, y:%v \n", radius, x, y)
}

type GreenCircle struct {
}

func (*GreenCircle) drawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle [color: green, radius: %v, x: %v, y:%v \n", radius, x, y)
}

type Shape interface {
	draw()
}

type Circle struct {
	x, y, radius int
	drawAPI      DrawAPI
}

func (c *Circle) draw() {
	c.drawAPI.drawCircle(c.radius, c.x, c.y)
}

func NewCircle(x, y, radius int, drawAPI DrawAPI) Shape {
	circle := &Circle{x, y, radius, drawAPI}
	return circle
}
