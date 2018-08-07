package main

import "fmt"

func main() {
	// finger := 4 - can be declared before or inside of switch
	switch finger := 4; finger { // scope of finger is limited to switch block
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Printf("Middle")
	case 4:
		fmt.Printf("Ring")
	case 5:
		fmt.Printf("Pinky")
	case 6:
		fmt.Printf("hello Hannibal")
	case 7,8:
		fmt.Printf("yeah, right")
	default:
		fmt.Printf("get out of here you freak")
	}
	fmt.Printf("\n")

	// - duplicated cases are not allowed
	// - expressionless switch:
	//	- expression in switch is optional, and in this case switch is true
	//	- expressions can be defined for each 'case' statement

	// fallthrough - to transfer control to the 1st case statement after the one that has been executed.
	// it should be the last statement in a 'case'. otherwise, compiler will complain
	switch num := 60; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}
}
