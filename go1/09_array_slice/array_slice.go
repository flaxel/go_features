package main

import (
	"fmt"
)

func main() {
	// ARRAYS - collection of elements with the same type, cannot be resized
	// value types: assigned to a new variables -> copy of the original array
	var a [3]int
	fmt.Println(a)

	b := [3]int{12, 78}
	fmt.Println(b)

	b[2] = 42
	fmt.Println(b)

	c := []int{1, 2, 3, 4}
	fmt.Println(c)

	fmt.Printf("|a|: %d, |b|: %d, |c|: %d\n", len(a), len(b), len(c))

	s := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary, compiler will complain if you omit this comma
	}
	fmt.Println(s)

	// SLICES - convenient/flexible/powerful wrapper on top of an array, references to existing arrays, dynamic
	// creation with composite literals or make(type, capacity)
	slice1 := c[1:4]
	fmt.Println(slice1)

	slice2 := []int{6, 7, 8}
	fmt.Println(slice2)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array init", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	// array: 7, slice created from index 1, capacity of slice is the number of elements in array starting from index 1
	// [...] - compiler count the array elements
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fmt.Println(fruitarray)
	fruitslice := fruitarray[1:3]
	fmt.Println(fruitslice)
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)] //re-slicing slice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
	fmt.Println(fruitslice)

	// create a slice by passing the type, length and capacity
	slice3 := make([]int, 5, 5)
	fmt.Println(slice3)
	slice3 = append(slice3, 3) // variadic function - accept variable number of arguments, create a new array
	fmt.Println(slice3)

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	// ATTENTION
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtractOne(nos)
	fmt.Println("slice after function call", nos)

	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	fmt.Println(pls)

	// memory optimisation - slices hold a reference to an array, as long as the slice is in memory, array cannot be garbage collected
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func subtractOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	// copy o fthe slice, using new slice, original array can be garbage collected
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2] // creates a slice of countries barring the last 2 elements
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
