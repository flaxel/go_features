package main

import (
	"go1/06_package/rectangle"
	"fmt"
	_ "math" // error silencer
)

// cannot be called explicitly, perform initialisation tasks
func init() {
	fmt.Println("Initializing 1 ...")
}

func init() {
	fmt.Println("Initializing 2 ...")
}

func main() {
	var length, width float64 = 6, 7
	fmt.Println("Geometrical Shape Properties")
	fmt.Println("Area:", rectangle.Area(length, width))
	fmt.Println("Diagonal:", rectangle.Diagonal(length, width))
}
