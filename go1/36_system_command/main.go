package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// source: https://golangcode.com/execute-a-command/
	cmdOutput, err := exec.Command("go", "version").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", cmdOutput)

	// COMMAND LINE ARGUMENTS
	// example command:
	// - go run main.go test
	// - go run main.go -num=22
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	// COMMAND FLAGS
	num := flag.Int("num", 42, "one number")
	flag.Parse()

	fmt.Println("num: ", *num)
	fmt.Println("tail: ", flag.Args())

	// ENVIRONMENT VARIABLES
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
