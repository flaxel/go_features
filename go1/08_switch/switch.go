package main

import (
	"fmt"
)

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}

	return false
}

func main() {
	// case is executed - leave the switch statement
	// finger := 4
	// switch finger {
	switch finger := 4; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	case 6, 7, 8, 9, 10:
		fmt.Println("this is your second hand")
	default:
		fmt.Println("incorrect finger number")
	}

	// FALLTHROUGH
	switch num := 15 * 5; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		//fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
		fallthrough
	default:
		fmt.Println("test")
	}

	// TYPE SWITCH
	var t interface{}
	t = true
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}
