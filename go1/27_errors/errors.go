package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// errors indicate an abnormal condition in the program
// errors in go are plain old values
// errors are represented using the built-in error type
// if a function/method returns an error then by convention it has to be the last value returned from the function
// idiomatic way: compare the returned error to nil
// error type with method: Error() string
// create a new error with: errors.New(info) or fmt.Errorf(info)
func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	// ASSERT STRUCT TYPE + EXTRACT MORE INFORMATION FROM THE STRUCT FIELDS
	f2, err2 := os.Open("test.txt")
	if err2, ok := err2.(*os.PathError); ok {
		fmt.Println("File at path", err2.Path, "failed to open")
	} else {
		fmt.Println(f2.Name(), "opened successfully")
	}

	// ASSERT STRUCT TYPE + EXTRACT MORE INFORMATION USING METHODS
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error:", err)
		}
	} else {
		fmt.Println(addr)
	}

	// DIRECT COMPARISON
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)

	// => NEVER EVER IGNORE AN ERROR
}
