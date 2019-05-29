package main

import (
	"fmt"
)

func main() {
	// no sorting, reference type, creation also with composite literals
	// SET - default value is nil
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println(personSalary)

	personSalary = map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	fmt.Println(personSalary)

	// GET
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])

	employee = "jamie2"
	fmt.Println("Salary of", employee, "is", personSalary[employee]) // return the default value for the type

	// newEmployee := "jamie"
	newEmployee := "joe"
	v, ok := personSalary[newEmployee]
	if ok {
		fmt.Println("Salary of", newEmployee, "is", v)
	} else {
		fmt.Println(newEmployee, "not found")
	}

	// ITERATION
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	// DELETE
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve2")
	// delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)

	// LENGTH
	fmt.Println("length is", len(personSalary))
}
