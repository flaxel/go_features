package main

import (
	"fmt"
	"runtime"
	"time"
)

// source: https://golangcode.com/print-the-current-memory-usage/
func main() {
	// starting memory usage
	printMemoryUsage()

	var overall [][]int

	for i := 0; i < 4; i++ {
		// allocate memory + append to overall => it does not garbage collected
		a := make([]int, 0, 999999)
		overall = append(overall, a)

		printMemoryUsage()
		time.Sleep(time.Second)
	}

	// clear memory
	overall = nil
	printMemoryUsage()

	// force gc => memory drop
	runtime.GC()
	printMemoryUsage()
}

func printMemoryUsage() {
	status := new(runtime.MemStats)
	runtime.ReadMemStats(status)

	// var status runtime.MemStats
	// runtime.ReadMemStats(&status)

	fmt.Printf("Alloc = %v MB", bToMb(status.Alloc))
	fmt.Printf("\tTotalAlloc = %v MB", bToMb(status.TotalAlloc))
	fmt.Printf("\tSys = %v MB", bToMb(status.Sys))
	fmt.Printf("\tNumGC = %v\n", status.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
