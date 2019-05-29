package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

// Channels: can be thought as pipes using which Goroutines communicate
// has a type associated with it, type of data that hte channel is allowed to transport
// zero value: nil
// read from channel: data := <- channel
// write to channel: channel <- data
// sends and receives are blocking by default:
// 	- data is sent to a channel -> control is blocked in the send statement until some other Goroutines reads from the Channel
//	- data is read from a channel, the read is blocked until some Goroutine writes data to that channel
// buffer size is 0 - unbuffered or synchronous channel
func main() {
	done := make(chan bool)
	go hello(done)
	<-done // blocking main goroutine, does not store the received data
	fmt.Println("main function")

	// ANOTHER EXAMPLE
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	// UNIDIRECTIONAL CHANNELS: send or receive data
	// sendch := make(chan<- int)
	// go sendData(sendch)
	// fmt.Println(<-sendch)

	// possible to convert a bidirectional channel to send only or receive only channel but not the vice versa
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)

	// CLOSING CHANNELS + FOR LOOPS: to notify receivers that no more data will be sent on the channel
	ch := make(chan int)
	go producer(ch)
	/* for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}*/
	for v := range ch {
		fmt.Println("Received ", v)
	}

	// REFACTOR CALC_SQUARES AND CALC_CUBES
	number2 := 589
	sqrch2 := make(chan int)
	cubech2 := make(chan int)
	go calcSquares2(number2, sqrch2)
	go calcCubes2(number2, cubech2)
	squares2, cubes2 := <-sqrch2, <-cubech2
	fmt.Println("Final output", squares2+cubes2)

	// DEADLOCK: sending data on a channel -> expected that some other Goroutine should receiving the data
	chDead := make(chan<- int)
	chDead <- 5
}
