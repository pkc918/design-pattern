package main

import "testing"

func TestShapeFactory_GetShape(t *testing.T) {
	producer := &FactoryProducer{}
	shapeFactory := producer.GetFactory("SHAPE")
	shapeFactory.GetShape("RECTANGLE").Draw()
	shapeFactory.GetShape("CIRCLE").Draw()
	shapeFactory.GetShape("SQUARE").Draw()
}

func TestColorFactory_GetColor(t *testing.T) {
	producer := &FactoryProducer{}
	shapeFactory := producer.GetFactory("COLOR")
	shapeFactory.GetColor("RED").Fill()
	shapeFactory.GetColor("GREEN").Fill()
	shapeFactory.GetColor("BLUE").Fill()
}
