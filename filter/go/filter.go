package main

import (
	"strings"
)

type Person struct {
	name          string
	gender        string
	maritalStatus string
}

func NewPerson(name, gender, maritalStatus string) Person {
	return Person{name: name, gender: gender, maritalStatus: maritalStatus}
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) getGender() string {
	return p.gender
}

func (p *Person) getMaritalStatus() string {
	return p.maritalStatus
}

type Criteria interface {
	meetCriteria(persons []Person) []Person
}

type CriteriaMale struct {
}

func (c *CriteriaMale) meetCriteria(persons []Person) []Person {
	malePersons := make([]Person, 0, 5)
	for _, person := range persons {
		if strings.ToUpper(person.getGender()) == "MALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type CriteriaFemale struct {
}

func (c *CriteriaFemale) meetCriteria(persons []Person) []Person {
	femalePersons := make([]Person, 0, 5)
	for _, person := range persons {
		if strings.ToUpper(person.getGender()) == "FEMALE" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {
}

func (c *CriteriaSingle) meetCriteria(persons []Person) []Person {
	singlePersons := make([]Person, 0, 5)
	for _, person := range persons {
		if strings.ToUpper(person.getMaritalStatus()) == "SINGLE" {
			singlePersons = append(singlePersons, person)
		}
	}
	return singlePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewAndCriteria(criteria, otherCriteria Criteria) *AndCriteria {
	return &AndCriteria{criteria: criteria, otherCriteria: otherCriteria}
}

func (c *AndCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaPersons := c.criteria.meetCriteria(persons)
	return c.otherCriteria.meetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewOrCriteria(criteria, otherCriteria Criteria) *OrCriteria {
	return &OrCriteria{criteria: criteria, otherCriteria: otherCriteria}
}

func (c *OrCriteria) meetCriteria(persons []Person) []Person {
	firstCriteriaItems := c.criteria.meetCriteria(persons)
	otherCriteriaItems := c.otherCriteria.meetCriteria(persons)

	for _, person := range otherCriteriaItems {
		if !contains(firstCriteriaItems, person) {
			firstCriteriaItems = append(firstCriteriaItems, person)
		}
	}

	return firstCriteriaItems
}

func contains(slice []Person, target Person) bool {
	for _, value := range slice {
		if value == target {
			return true // 如果找到目标对象，则返回true
		}
	}
	return false // 如果遍历完切片都没有找到目标对象，则返回false
}
