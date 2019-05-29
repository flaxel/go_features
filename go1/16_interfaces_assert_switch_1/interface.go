package main

import (
	"fmt"
)

type SalaryCalculator interface {
	Calculate() int
}

type Permanent struct {
	empID    int
	baiscPay int
	pf       int
}

type Contract struct {
	empdID   int
	basicPay int
}

func (p Permanent) Calculate() int {
	return p.baiscPay + p.pf
}

func (c Contract) Calculate() int {
	return c.basicPay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0

	for _, v := range s {
		expense += v.Calculate()
	}

	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

type Tester interface {
	Test()
}

type MyFloat float64

func (m MyFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

// all types implement the empty interface
func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", v)
	case int:
		fmt.Printf("I am an integer and my value is %d\n", v)
	case Contract:
		fmt.Printf("I am a contract and must pay %d\n", v.Calculate())
	default:
		fmt.Printf("Unknown type\n")
	}
}

// interface: defines the behaviour of an object
// no code changes necessary if another employee wil be add
func main() {
	perm1 := Permanent{1, 5000, 20}
	perm2 := Permanent{2, 6000, 30}

	cemp1 := Contract{3, 3000}

	employees := []SalaryCalculator{perm1, perm2, cemp1}
	totalExpense(employees)

	// interface can be thought of as being represented internally by a tuple (type, value)
	var t Tester
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Test()

	// empty interface: zero methods
	s := "Hello World"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe2(strt)

	// Type Assertion: extract the underlying value of the interface
	// syntax: variable.(type)
	var s2 interface{} = 56
	assert(s2)

	var s3 interface{} = "test"
	assert(s3)

	// Type Switch
	findType("Naveen")
	findType(77)
	findType(89.98)
	findType(cemp1)
}
