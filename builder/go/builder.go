package main

import "fmt"

type Item interface {
	Name() string
	Packing() Packing
	Price() float64
}

type Packing interface {
	Pack() string
}

/* Packing */

type Wrapper struct {
}

func (*Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct {
}

func (*Bottle) Pack() string {
	return ""
}

// -----

/* Item */

type Burger struct {
}

func (*Burger) Packing() Packing {
	return &Wrapper{}
}

func (*Burger) Price() float64 {
	return 24.0
}

type ColdDrink struct {
}

func (*ColdDrink) Packing() Packing {
	return &Bottle{}
}

func (*ColdDrink) Price() float64 {
	return 15.0
}

// -----

/*Burger ColdDrink 扩展*/

type VegBurger struct {
	Burger
}
type ChickenBurger struct {
	Burger
}
type Coke struct {
	Burger
}
type Pepsi struct {
	Burger
}

func (v *VegBurger) Price() float64 {
	return 25.0
}

func (v *VegBurger) Name() string {
	return "Veg Burger"
}

func (v *ChickenBurger) Price() float64 {
	return 50.5
}

func (v *ChickenBurger) Name() string {
	return "Chicken Burger"
}

func (v *Coke) Price() float64 {
	return 30.0
}

func (v *Coke) Name() string {
	return "Coke"
}

func (v *Pepsi) Price() float64 {
	return 35.0
}

func (v *Pepsi) Name() string {
	return "Pepsi"
}

/*Meal*/

type Meal struct {
	items []Item
}

func (meal *Meal) AddItem(item Item) {
	meal.items = append(meal.items, item)
}

func (meal *Meal) GetCost() float64 {
	var cost float64
	for _, item := range meal.items {
		cost += item.Price()
	}
	return cost
}

func (meal *Meal) ShowItems() {
	for _, item := range meal.items {
		fmt.Printf("Item: %v \n", item.Name())
		fmt.Printf("Packing: %v \n", item.Packing().Pack())
		fmt.Printf("Price: %v \n", item.Price())
	}
}

type MealBuilder struct {
}

func (*MealBuilder) PrepareVegMeal() Meal {
	meal := Meal{}
	meal.AddItem(&VegBurger{})
	meal.AddItem(&Coke{})
	return meal
}

func (*MealBuilder) PrepareNonVegMeal() Meal {
	meal := Meal{}
	meal.AddItem(&ChickenBurger{})
	meal.AddItem(&Pepsi{})
	return meal
}
