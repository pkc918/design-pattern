package main

import (
	"fmt"
	"testing"
)

func TestMealBuilder_PrepareVegMeal(t *testing.T) {
	mealBuilder := MealBuilder{}
	vegMeal := mealBuilder.PrepareVegMeal()
	fmt.Println("Veg Meal")
	vegMeal.ShowItems()
	fmt.Printf("Total Cost: %f \n", vegMeal.GetCost())
}

func TestMealBuilder_PrepareNonVegMeal(t *testing.T) {
	mealBuilder := MealBuilder{}
	nonVegMeal := mealBuilder.PrepareNonVegMeal()
	fmt.Println("Non-Veg Meal")
	nonVegMeal.ShowItems()
	fmt.Printf("Total Cost: %f \n", nonVegMeal.GetCost())
}
