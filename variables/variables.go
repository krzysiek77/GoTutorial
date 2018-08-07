package main

import (
	"fmt"
	"math"
)

func main() {
	var age int //variable declaration. default value = 0
	fmt.Println("my age is = ", age)
	age = 29
	fmt.Println("my age is = ", age)
	age = 54
	fmt.Println("my age is = ", age)

	var age2 int = 49
	fmt.Println("my age2 is = ", age2)

	//if var is initialized during declaration
	//Go will automatically infer the typo of that var
	var age3 = 39
	fmt.Println("my age is = ", age3)

	//multiple variable declarations
	var width, height int = 100, 50 // int can be dropped
	fmt.Println("width is ", width, "height is", height)

	//multiple variable declaration - without initialization
	var width2, height2 int //both are int
	fmt.Println("width is ", width2, "height is", height2)
	width2 = 100
	height2 = 50
	fmt.Println("width is ", width2, "height is", height2)

	//multiptle variables of different types
	var (
		name = "Aneta"
		age4 = 29
		height4 int
	)
	fmt.Println("my name is ", name, ", age is ", age4, " and height is ", height4)

	//'short hand' declaration (:=)
	//it requires initial values for all variable in the left hand side of the assignment
	name5, age5 := "John", 33
	fmt.Println(name5, " is of age ", age5)

	//can be used only when at least one of the variables in the left side of := is newly created
	a, b := 20, 30
	fmt.Println("a = ", a, ", b = ", b)
	b, c := 40, 50
	fmt.Println("b = ", b, ", c = ", c)
	b, c = 60, 70
	fmt.Println("b = ", b, ", c = ", c)

	//variables can also be assigned values which are computed during run time
	x, y := 123.2, 233.3
	z := math.Min(x, y)
	fmt.Println("minimum value is ", z)

	//Go is strongly type, which means variables declared as belonging to one type
	//cannot be assigned a value of another type
	//age6 := 29
	//age6 = "John" //error
}
