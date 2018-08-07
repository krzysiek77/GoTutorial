package main

import (
	"fmt"
	"unicode/utf8"
)

// string
// 	- is a slice of bytes
//	- is Unicode compliant
//	- UTF-8 Encoded (In UTF-8 encoding a code point can occupy more than 1 byte)
//	- immutable; once created, cannot be changed. To change a string
//		- convert it to slice of runes
//		- make a change
//		- return as a new string using: string(s)

func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func printChars(s string)  {
	runes := []rune(s)
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

// rune:
// 	- builtin type in Go and it's an alias of int32
//	- it represent Unicode code point
//	- any valid unicode char within single quotes
func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func main() {
	name := "Hello World"
	printBytes(name)
	fmt.Println("")
	printChars(name)
	fmt.Println("")

	name2 := "Who åre you?" // å - occupies 2 bytes
	printBytes(name2)
	fmt.Println("")
	printChars(name2)
	fmt.Println("")
	printCharsAndBytes(name2)

	stringFromBytesSlice()

	// length
	length(name)
	length(name2)

}

func stringFromBytesSlice() {
	hexByteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(hexByteSlice)
	fmt.Println(str)

	decByteSlice := []byte{67, 97, 102, 195, 169}
	str = string(decByteSlice)
	fmt.Println(str)

	runSlice := []rune{0x0043, 0x0061, 0x0066, 0x00E9}
	str = string(runSlice)
	fmt.Println(str)
}

// length of string - number of runes in it
func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}