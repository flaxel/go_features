package main

import "fmt"
import "sort"

// implement sort.Interface
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// changes the given slice and doesn't return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// CUSTOM COMPARATOR
	sport := []string{"football", "baseball", "rocket"}
	sort.SliceStable(sport, func(i, j int) bool {
		return len(sport[i]) < len(sport[j])
	})
	fmt.Println(sport)

	// SORT BY FUNCTION
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
	sort.Sort(sort.Reverse(byLength(fruits)))
	fmt.Println(fruits)
}
