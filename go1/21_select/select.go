package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

// used to choose from multiple send/receive channel operations
// blocks until one of the send/receive operation is ready
// i fmultiple operations are ready, one of them is chosen at random
// syntax similar to switch expcept that each of the case statement will be a channel operation
// return the quickest response to the user
func main() {
	// EXAMPLE 1
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	// DEFAULT CASE AND DEADLOCK
	ch3 := make(chan string)
	select {
	case <-ch3:
	default: // without default, you get a deadlock
		fmt.Println("default case executed")
	}

	// EMPTY SELECT
	// select {} // statement will block until one of its cases is executed -> deadlock

	// EXAMPLE 2 -
	ch := make(chan string)
	// var ch chan string
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		case <-time.After(50 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}
