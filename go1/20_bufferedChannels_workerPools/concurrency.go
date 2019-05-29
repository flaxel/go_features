package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() // decrement the counter value
}

var jobs = make(chan Job, 10)       // listen for new tasks
var results = make(chan Result, 10) // write result to buffered channel

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// Create a worker Goroutine
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// Create a pool of worker Goroutines
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

// sends to a buffered channel are blocked only when the buffer is full
// receives from a buffered channel are blocked only when the buffer is empty
// create a buffered channel: channel := make (chan type, capacity), capacity > 0 to have an buffer
// using like a semaphore
func main() {
	// EXAMPLE 1
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// EXAMPLE 2
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v, "from ch2")
		time.Sleep(2 * time.Second)

	}

	// LENGTH VS CAPACITY
	// length: number of elements currently queued in it
	// capacity: number of values that the channel can hold
	ch3 := make(chan string, 3)
	ch3 <- "naveen"
	ch3 <- "paul"
	fmt.Println("capacity is", cap(ch3))
	fmt.Println("length is", len(ch3))
	fmt.Println("read value", <-ch3)
	fmt.Println("capacity is", cap(ch3))
	fmt.Println("new length is", len(ch3))

	// WAIT GROUP: used to wait for a collection of Goroutines to finish executing, control is blocked until all Goroutines are  are finished
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)          // increment counter value
		go process(i, &wg) // must be the address, otherwise each goroutine will have its onw copy of the WaitGroup
	}
	wg.Wait() // blocks the Goroutine until the counter becomes zero
	fmt.Println("All go routines finished executing")

	// WORKER POOL
	startTime := time.Now()

	noOfJobs := 100
	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)

	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}
