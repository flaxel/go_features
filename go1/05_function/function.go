package main

import (
	"fmt"
)

/*
general syntax:

func <name>(<parameter_name> <type>) <return_type> {
	<body>
}

parameters and return type are optional
*/
// see also 30_first_class_functions
func main() {
	fmt.Println("Total price is", calculateBill(10, 5))
	fmt.Println()

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area: %f, Perimeter: %.2f\n", area, perimeter)

	area, _ = rectProps(11.4, 3.2) // perimeter is discarded - using the blank identifier
	fmt.Printf("Area: %f, Perimeter: %.2f\n", area, perimeter)
	fmt.Println()

	area2, perimeter2 := rectProps2(10.8, 5.6)
	fmt.Printf("Area2: %f, Perimeter2: %.2f\n", area2, perimeter2)
	fmt.Println()

	// RECURSION
	fmt.Println(fact(7))
}

// func calculateBill(price int, number int) int {
func calculateBill(price, number int) int {
	totalPrice := price * number
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
