package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrementMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incrementChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

// critical section: modify shared resources
// mutex: provide a locking mechanism to ensure that only one Goroutine is running the critical section of code, prevent race condition
func main() {
	// SOLVING RACE CONDITIONS WITH MUTEX
	var w sync.WaitGroup
	var m sync.Mutex // zero value, struct type

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementMutex(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)

	// SOLVING RACE CONDITIONS WITH CHANNELS
	ch := make(chan bool, 1) // buffered channel with capacity 1, use write blocking
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementChannel(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)

	// MUTEX VS CHANNELS
	// there is no wrong in choosing either
	// answer lies in the problem you are trying to solve
	// use channels when Goroutines need to communicate with each other
	// use mutexes when only one Goroutine should access the critical section code
}
