package main

import (
	"math"
	"fmt"
)

func main() {
	// value of a constant should be known at compile time.
	var a = math.Sqrt(4) //allowed
	//const b = math.Sqrt(9) //not allowed, because it's evaluated at the run time

	// string constant are untyped - does not have a type
	// since Go is strongly typed language, untyped constants have a default type
	// associated with them and they supply it if and only if a line of code demands it
	
	// it's possible to create a typed constant
	const c string = "Hello World" //it's a constant of type string

	// mixing types during assignment is not allowed
	//var x = "Sam"
	//type myString string
	//var y myString = "Sam" //since constant 'Sam' is untyped it can be assigned to any string variable
	//x = y //not allowed

	// boolean constants behave the same way as string

	// numeric constants
	// type of each variable is determined by the syntax of the numeric constant
	// z is generic (can represent a float, int, etc) and hence it is possible to be assigned to any compatible type
	const z = 5 //untyped of value 5
	var intVar int = z
	var int32Var int32 = z
	var float64Var float64 = z
	var complex64Var = z
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

}
