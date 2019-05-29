package main

import (
	"fmt"
)

func main() {
	// functionality: converting the variable number of arguments to a new slice
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	// SLICE
	slice := []int{1, 2, 3}
	find(2, slice...) // wrapping a slice to a slice is not possible - slice is directly passed to the function without a new slice being created

	// MODIFYING
	welcome := []string{"hello", "world"}
	change(welcome...) // there is no copy and no new slice is created
	fmt.Println(welcome)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Println()
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground") // create a new slice
	fmt.Println(s)
}
