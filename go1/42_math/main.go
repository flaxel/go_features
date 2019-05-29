package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// default number generator is deterministic
	// Intn: n >= 0 und n < 100
	// Float64: n >= 0 und n < 1
	fmt.Println(rand.Intn(42))
	fmt.Println(rand.Intn(100))
	fmt.Println(randomInt(10, 20))

	fmt.Println(rand.Float64())
	fmt.Println(randomFloat(5, 10))

	// VARYING SEQUENCES - not safe to use for random numbers you intend to be secret, use crypto/rand package
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Float64())
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomFloat(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
