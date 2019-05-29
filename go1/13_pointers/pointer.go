package main

import (
	"fmt"
)

// pointer: variable which stores the memory address of another variable
func main() {
	//DECLARING
	b := 255
	var a *int
	fmt.Println("pointer value before:", a)
	a = &b
	// a := &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// DEREFERENCING
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)

	// FUNCTIONS
	change(a)
	fmt.Println("value of a after function call is", b)

	// ATTENTION: DO NOT PASS A POINTER TO AN ARRAY AS A ARGUMENT TO A FUNCTION. USE SLICE INSTEAD.
	// a[x] is shorthand for (*a)[x]. But it is better to use slices
	array := [3]int{89, 90, 91}
	modify(&array)
	fmt.Println(array)

	array2 := [3]int{89, 90, 91}
	modify2(array2[:])
	fmt.Println(array2)

	// NOTICES:
	// NO support for pointer arithmetic
}

func change(val *int) {
	*val = 55
}

func modify(arr *[3]int) {
	arr[0] = 90
}

func modify2(sls []int) {
	sls[0] = 90
}
