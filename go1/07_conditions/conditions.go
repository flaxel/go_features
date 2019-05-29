package main

import (
	"fmt"
)

func main() {
	// IF - ELSE - STATEMENT
	// num := 10
	// if num%2 == 0 {
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	fmt.Println()

	// LOOP - all components are optional
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}

		if i == 3 {
			continue
		}

		fmt.Printf("%02d ", i)
	}
	fmt.Printf("\n\n")

	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Printf("\n\n")

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
	fmt.Println()

	// NESTED LOOPS + BREAK
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
	fmt.Println()

	// RANGE
	a := []float64{67.7, 89.8, 21, 78}
	sum := 0.0
	for i, v := range a { // using blank identifier ('_') to ignore index
		fmt.Printf("%d.element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Printf("\nsum of all elements of a: %f\n", sum)

	// INFINITE LOOPS
	/* for {
		fmt.Println("Hello World!")
	} */
}
