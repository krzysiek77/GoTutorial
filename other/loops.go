package main

import "fmt"

// go has only 'for' loop.
// all 3 components (init, cond, post) are optional. semicolons can be omitted as well
//
// for initialisation; condition; post {}
//
// additional statements: break, continue

// infinite loop:
// for {}

func main() {
	// print all even numbers

	i := 0
	for ; i <= 10; {
		fmt.Printf("%d ", i)
		i += 2
	}
}
