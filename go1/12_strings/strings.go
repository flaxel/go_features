package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// strings are UTF-8 encoded, slice of bytes, immutable
func main() {
	// ELEMENTS
	name := "Hello World!"
	printBytes(name)
	printChars(name)

	name = "Señor"
	printBytes(name)
	printChars(name)
	printCharsAndBytes(name)

	// BUILDING
	// byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	// byteSlice := []byte{67, 97, 102, 195, 169}
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	// str := string(byteSlice)
	str := string(runeSlice)
	fmt.Println(str)

	// LENGTH
	length(name)

	// MUTATING
	h := "hello"
	fmt.Println(mutate([]rune(h)))
	fmt.Println()

	// INFO
	info()

	// FORMATTING
	formatting()

	// NUMBER PARSING
	parsing()
}

// prints bytes in the hexadecimal format
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

// print characters of the string
// code point of 'ñ' is two bytes long but char is only one byte -  solve this problem with rune
// rune: alias of int32, represents unicode code point, does not matter how many bytes the code point occupies
func printChars(s string) {
	rune := []rune(s)
	for i := 0; i < len(rune); i++ {
		fmt.Printf("%c ", rune[i])
	}
	fmt.Println()
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	s[1] = 'a'
	return string(s)
}

// https://golang.org/pkg/strings/
func info() {
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	fmt.Println()
}

type point struct {
	x, y int
}

func formatting() {
	p := point{1, 2}

	// print instance
	fmt.Printf("%v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.
	fmt.Printf("%+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.
	fmt.Printf("%#v\n", p)

	// To print the type of a value, use `%T`.
	fmt.Printf("%T\n", p)

	// Formatting booleans is straight-forward.
	fmt.Printf("%t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.
	fmt.Printf("%d\n", 123)

	// This prints a binary representation.
	fmt.Printf("%b\n", 14)

	// This prints the character corresponding to the
	// given integer.
	fmt.Printf("%c\n", 33)

	// `%x` provides hex encoding.
	fmt.Printf("%x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.
	fmt.Printf("%f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// For basic string printing use `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("%x\n", "hex this")

	// To print a representation of a pointer, use `%p`.
	fmt.Printf("%p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.
	fmt.Fprintf(os.Stderr, "an %s\n\n", "error")
}

func parsing() {
	// bits of precision: 64
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `0` means infer the base from the string
	// `64` requires that the result fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// recognize hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// parsing uint
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// convenience function for basic base-10 `int` parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
