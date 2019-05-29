package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World Goroutine!")
}

// go is a concurrent langauge and not a parallel one
// concurrency: capability to deal with lots of things at once
// parallelism: capability to deal with lots of things at the same time
// concurrency is handled using Goroutines and channels
// Goroutine: function or method that run concurrently with other functions or methods, light weight threads
// Advantages:
// 	- extremly cheap when compared to threads
// 	- multiplexed to fewer number of OS threads, one thread with 1.000 of Goroutines, thread blocks waiting for user input
// 		-> another OS thread is created and remaining Goroutines are moved to the new OS thread
// 	- communicate using channels, prevent race conditionsfrom happening when accessing shared memory using Goroutines
// create a Goroutine with keyword 'go'
func main() { // runs in the main Goroutine
	// create new Goroutine - call return immediately -> does not wait for the Goroutine to finishing executing
	// next line of code + any return values from the Goroutine are ignored
	// main Goroutine termiantes -> programm will be terminated
	go hello()
	time.Sleep(1 * time.Second) // better using channels to block the main Goroutine until all other Goroutines finish their execution
	fmt.Println("main function")
}
