package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Person2 struct {
	string
	int
}

type Address struct {
	city, state string
}

type Person struct {
	name    string
	age     int
	address Address
}

type PersonPromoted struct {
	name string
	age  int
	Address
}

type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}

func main() {
	// with type
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	// creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)

	// annonymous structure, without a new type
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	// ZERO
	var emp4 Employee // zero valued structure, properties with default value ("" and 0)
	fmt.Println("Employee 4", emp4)

	// ACCESS
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)

	// POINTERS
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	// ANNONYMOUS FIELDS - names are the type names
	p2 := Person2{"Naveen", 50}
	fmt.Println(p2)

	//NESTED
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println(p)
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	// PROMOTED FIELDS
	var pp PersonPromoted
	pp.name = "Naveen"
	pp.age = 50
	pp.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println(pp)
	fmt.Println("Name:", pp.name)
	fmt.Println("Age:", pp.age)
	fmt.Println("City:", pp.city)   //city is promoted field
	fmt.Println("State:", pp.state) //state is promoted field

	// STRUCTS EQUALITY - struct variables are not comparable if they contain fields which are not comparable
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
}
