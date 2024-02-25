package main

type Cloneable interface {
	Clone() Cloneable
}

type ShapeCache struct {
	Hashtable map[string]Cloneable
}

func (s *ShapeCache) GetShape(shapeId string) Cloneable {
	return s.Hashtable[shapeId].Clone()
}

func (s *ShapeCache) LoadCache(shapeId string, shape Cloneable) {
	s.Hashtable[shapeId] = shape
}
