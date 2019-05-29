package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// READ AN ENTIRE FILE INTO MEMORY
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println("file reading error", err)
		return
	}

	fmt.Println("Content of file:", string(data))

	// PROBLEM: BINARY IS INDEPENDENT OF THE SOURCE CODE + CAN BE RUN FROM ANY LOCATION + FILE NOT FOUND!
	//	- absolute file path
	//	- passing file path as command line flag
	// 	- bundling file along with the binary

	// ABSOLUTE FILE PATH: specification in the program
	data, err = ioutil.ReadFile("/home/flaxel/Dokumente/go/src/32_reading/test.txt")

	if err != nil {
		fmt.Println("file reading error", err)
		return
	}

	fmt.Println("Content of file:", string(data))

	// PASSING FILE PATH AS COMMAND LINE FLAG
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	data, err = ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Content of file:", string(data))

	// BUNDLING THE TEXT FILE ALONG WITH BINARY
	// using packr: converts static files to go files
	box := packr.NewBox(".")        // represents a folder whose contents will be embedded to the binary
	data2 := box.String("test.txt") // reading
	fmt.Println("Contents of file:", data2)

	// bundle file to binary: packr install -v <folder>

	// READING A FILE IN SMALL CHUNKS: extremely large files - read in small chunks
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := bufio.NewReader(f) // buffered reader
	b := make([]byte, 3)    // byte slice with capacity 3

	for {
		_, err := r.Read(b)

		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}

		fmt.Println(string(b))
	}

	// READING FILE LINE BY LINE
	f, err = os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() { // read next line
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
