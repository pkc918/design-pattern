package main

import "testing"

var cache *ShapeCache

func init() {
	cache = &ShapeCache{Hashtable: make(map[string]Cloneable)}
	cache.LoadCache("Rectangle", &Rectangle{Type: "Rectangle"})
	cache.LoadCache("Circle", &Circle{Type: "Circle"})
}

type Rectangle struct {
	Type string
}

func (r *Rectangle) Clone() Cloneable {
	rClone := *r
	return &rClone
}

type Circle struct {
	Type string
}

func (c *Circle) Clone() Cloneable {
	cClone := *c
	return &cClone
}

func TestClone(t *testing.T) {
	p1 := cache.Hashtable["Rectangle"]
	p2 := p1.Clone()
	if p1 == p2 {
		t.Fatal("get clone not working")
	}
}

func TestCloneFromCache(t *testing.T) {
	p := cache.GetShape("Circle").Clone()
	t1 := p.(*Circle)
	if t1.Type != "Circle" {
		t.Fatal("error")
	}
}
