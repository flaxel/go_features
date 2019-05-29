package main

import "fmt"

// x := "test" // not possible outside a body
var x = "test"

func main() {
	var age int
	fmt.Println("my age is ", age)
	age = 42
	fmt.Println("my age is ", age)

	var age2 = 54 // type will be inferred
	fmt.Println("my age is ", age2)

	var (
		name   = "naveen"
		age3   = 29
		height int
	)
	fmt.Println("my name is", name, ", age is", age3, "and height is", height)

	name2 := "test" //short hand declaration
	fmt.Println(name2)

	a, b := 42, 43 // only if one variable at the left side is new
	fmt.Println("a: ", a, "b: ", b)
	// a,b:= 11,12 // error no new variable
	b, c := 23, 24
	fmt.Println("a: ", a, "b: ", b, "c: ", c)

	// shorthand for defining variables in a bulk
	var (
		d, e string = "test D", "test E"
		f, g int    = 8, 9
	)

	fmt.Printf("d: %s, e: %s, f: %d, g: %d\n", d, e, f, g)
}
