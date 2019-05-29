package main

import (
	"fmt"
	"math"
)

func main() {
	// NEW VALUE
	const c = 42
	const d float64 = 45.23
	fmt.Println("c:", c, "d:", d)
	// c = 2 // reassignment not allowed
	fmt.Println()

	// COMPILE TIME
	s := math.Sqrt(4)
	// const b = math.Sqrt(4) // not allowed - constant should be known at compile time
	fmt.Println("sqrt:", s)
	fmt.Println()

	// STRING
	const hello = "hello" // untyped
	fmt.Printf("%T: %s\n", hello, hello)
	fmt.Println()

	// NUMERIC
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	fmt.Println()
}
