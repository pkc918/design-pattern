package main

import "testing"

func TestName(t *testing.T) {
	redCircle := NewCircle(100, 100, 10, &RedCircle{})
	greenCircle := NewCircle(100, 100, 10, &GreenCircle{})
	redCircle.draw()
	greenCircle.draw()
}
