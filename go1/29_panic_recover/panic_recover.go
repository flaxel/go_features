package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		debug.PrintStack()
	}
}

func fullNameRecover(firstName *string, lastName *string) {
	defer recovery()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func aRuntime() {
	defer recovery()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

// situations where the program cannot simply continue executing after an abnormal situation
// panic: terminate the program

// function encounters a panic -> execution is stopped + any deferred functions are executed + control return to its caller,
// process continues until all functions of the current Goroutine have returned at which point the program prints the panic message,
// followed by the stack trace and then terminates

// recover: used to regain control of a panicking program
// recover is useful only when called inside deferred functions -> stops the panicking sequence by restoring normal execution + retrieves the error value passed to the call of panic
// if recover is called outside the deferred dunction, it will not stop a panicking sequence

// panic + recover can be considered similar to try-catch-finally idiom

// ATTENTION: SHOULD AVOID PANIC AND RECOVER AND USE ERRORS

// possible use cases:
// 	- an unrecoverable error where the program cannot simply continue its execution, for example: binding to server port fails
// 	- a programmer error, for example: use nil as pointer
func main() {
	// PANIC
	/* defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main") */

	// RECOVER
	/* defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullNameRecover(&firstName, nil)
	fmt.Println("returned normally from main") */

	// PANIC, RECOVER, GOROUTINES
	// panic will not recovered -> recovery function is present in a different goroutine
	/* a()
	fmt.Println("normally returned from main") */

	// RUNTIME PANICS
	aRuntime()
	fmt.Println("normally returned from main")
}
