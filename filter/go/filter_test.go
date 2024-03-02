package main

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	persons := make([]Person, 0, 5)

	persons = append(persons, NewPerson("Robert", "Male", "Single"))
	persons = append(persons, NewPerson("John", "Male", "Married"))
	persons = append(persons, NewPerson("Laura", "Female", "Married"))
	persons = append(persons, NewPerson("Diana", "Female", "Single"))
	persons = append(persons, NewPerson("Mike", "Male", "Single"))
	persons = append(persons, NewPerson("Bobby", "Male", "Single"))

	male := &CriteriaMale{}
	female := &CriteriaFemale{}
	single := &CriteriaSingle{}
	singleMale := NewAndCriteria(single, male)
	singleOrFemale := NewOrCriteria(single, female)

	fmt.Println("Males: ")
	printPersons(male.meetCriteria(persons))

	fmt.Println("Females: ")
	printPersons(female.meetCriteria(persons))

	fmt.Println("Single Males: ")
	printPersons(singleMale.meetCriteria(persons))

	fmt.Println("Single Or Females: ")
	printPersons(singleOrFemale.meetCriteria(persons))
}

func printPersons(persons []Person) {
	for _, person := range persons {
		fmt.Printf("Person: [Name: %v, Gender: %v, Marital Status: %v]\n", person.getName(), person.getGender(), person.getMaritalStatus())
	}
}
