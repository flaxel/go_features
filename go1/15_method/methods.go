package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

type Employee2 struct {
	name string
	age  int
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// displaySalary() method has Employee as the receiver type
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/*
Method with value receiver
*/
func (e Employee2) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() { // TOP!
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() { // TOP!
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

/*
just a function with a special receiver type that is written between the func keyword and the method name

func (t Type) methodName(parameter list) {
}
*/
func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()

	/*
		WHY METHODS AND FUNCTIONS?
			- not a pure oo programming language, no support for classes
			- methods with the same name can be defined on different types whereas functions with the same names are not allowed
			- BUT NOT MIXTURE OF BOTH
	*/

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f\n", c.Area())

	// POINTER RECEIVERS VS VALUE RECEIVERS
	e := Employee2{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	(&e).changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//(&e).changeAge(51)
	e.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)

	/*
		RESULT:
			- pointer receivers: changes made to the receiver inside the method should be visible to the caller or it's expensive to copy a data structure
			- in all other situations - value receivers can be used
			- choose only one receiver for one project
	*/

	// ANNONYMOUS FIELDS
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress() // accessing fullAddress method of address struct

	// VALUE RECEIVERS IN METHODS VS VALUE ARGUMENTS IN FUNCTIONS
	r2 := rectangle{
		length: 10,
		width:  5,
	}
	area(r2)
	r2.area()

	p2 := &r2
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	// area(p)

	p2.area() // calling value receiver with a pointer

	/*
		RESULT:
			- When a function has a value argument, it will accept only a value argument.
			- When a method has a value receiver, it will accept both pointer and value receivers.
	*/

	// POINTER RECEIVERS IN METHODS VS POINTER ARGUMENTS IN FUNCTIONS

	p3 := &r2 //pointer to r
	perimeter(p3)
	p3.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	// perimeter(r)

	r2.perimeter() // calling pointer receiver with a value

	/*
		RESULT: functions with pointer arguments will accept only pointers whereas methods with pointer receivers will accept both value and pointer receiver
	*/

	// METHODS ON NON STRUCT TYPES
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
