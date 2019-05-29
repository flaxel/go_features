package main

import (
	"fmt"
)

type add func(a int, b int) int

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

// first class functions: allows functions to be assigned to variables, passed as arguments to other functions and returned from other functions
func main() {
	// ANONYMOUS FUNCTIONS: do not have a name
	a := func() {
		fmt.Println("hello world first class function")
	}

	a()
	fmt.Printf("%T\n", a)

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

	// USED DEFINED FUNCTION TYPES
	var a2 add = func(a int, b int) int {
		return a + b
	}
	s := a2(5, 6)
	fmt.Println("Sum", s)

	// HIGHER-ORDER FUNCTIONS: takes one or more functions as arguments OR returns a function as its result
	// 1) PASSING FUNCTIONS AS ARGUMENTS TO OTHER FUNCTIONS
	f := func(a, b int) int {
		return a + b
	}

	simple(f)

	// 2) RETURNING FUNCTIONS FROM OTHER FUNCTIONS
	s2 := simple2()
	fmt.Println(s2(60, 7))

	// CLOSURES: special cases of annonymous functions, access the variables defined outside the body of the function
	a3 := 5
	func() {
		fmt.Println("a3 =", a3)
	}()

	// every closure is bound to its own surrounding variable
	a4 := appendStr()
	b4 := appendStr()
	fmt.Println(a4("World"))
	fmt.Println(b4("Everyone"))

	fmt.Println(a4("Gopher"))
	fmt.Println(b4("!"))

	// PRACTICAL USE
	student1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	student2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	students := []student{student1, student2}
	filtered := filter(students, func(element student) bool {
		if element.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(filtered)

	values := []int{5, 6, 7, 8, 9}
	r := iMap(values, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
