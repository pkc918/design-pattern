package main

import "fmt"

type Employee struct {
	name, dept   string
	salary       int
	subordinates []*Employee
}

func NewEmployee(name, dept string, sal int) *Employee {
	return &Employee{
		name:         name,
		dept:         dept,
		salary:       sal,
		subordinates: make([]*Employee, 0, 5),
	}
}

func (e *Employee) findIndex(employee *Employee) int {
	for i, subordinate := range e.subordinates {
		if &subordinate == &employee {
			return i
		}
	}
	return -1
}

func (e *Employee) add(employee *Employee) {
	e.subordinates = append(e.subordinates, employee)
}

func (e *Employee) remove(employee *Employee) {
	i := e.findIndex(employee)
	fmt.Println(i)
	e.subordinates = append(e.subordinates[:i], e.subordinates[i+1:]...)
}

func (e *Employee) getSubordinates() []*Employee {
	return e.subordinates
}

func (e *Employee) toString() string {
	return fmt.Sprintf("Employee: [name: %v, dept: %v, salary: %v]\n", e.name, e.dept, e.salary)
}
