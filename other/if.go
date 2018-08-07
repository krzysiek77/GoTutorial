package main

import "fmt"

func main() {
	if 1 == 2 {

	} else if 2 ==2 {

	} else {

	}

	// other variant for if
	// if statement; condition { }
	// in case below, 'num' is available only in the if statement
	// 'else' statement needs to start in the same line after closing curly brace. otherwise, compiler error
	if num := 10; num % 2 == 0 {
		fmt.Println(num, " is even")
	} else {
		fmt.Println(num, " is odd")
	}
}
