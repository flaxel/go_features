package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// BOOLEAN
	a := true
	var b bool
	fmt.Println("a:", a, "b:", b)
	fmt.Println("a || b:", a || b)
	fmt.Println("a && b:", a && b)
	fmt.Println()

	// INTEGER: generally use int - only use another type if there is a need
	i := 42 // 32 or 64 bit number, depending on the operating system
	sizeOf := unsafe.Sizeof(i)
	fmt.Println("integer:", i)
	fmt.Printf("type of a is %T, size of a is %d\n", i, sizeOf)
	fmt.Println()

	// FLOAT
	f, g := 5.67, 8.97
	fmt.Printf("type of f %T g %T\n", f, g)
	sum := f + g
	diff := f - g
	fmt.Println("sum", sum, "diff", diff)
	fmt.Println()

	// COMPLEX
	c1 := 6 + 7i
	c2 := complex(6, 7)
	fmt.Println("c1:", c1, "c2:", c2)
	fmt.Println("sum:", c1+c2)
	fmt.Println("product:", c1*c2)
	fmt.Println()

	// MORE
	// byte - alias of uint8
	// rune - alias of int32

	// STRING
	first := "Naveen"
	var last string
	last = "Shermon"
	fmt.Println("My name is", first, last)
	fmt.Println()

	// CONVERSION
	d, e := 55, 67.8
	fmt.Println("Conversion sum:", d+int(e))
	fmt.Println("Conversion float:", float64(d))
	fmt.Println()

	// MIXING TYPES
	defaultName := "Sam"
	type myString string
	var customName myString = "Sam2"
	// customName = defaultName // not allowed
	fmt.Println("default:", defaultName, "custom:", customName)
}
