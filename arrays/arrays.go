package main

import (
	"fmt"
		)

// Array - collection of elements of the same type.
// Mixing elements of different types in array is not allowed in Go

func main() {
	var a [3]int //int array, length 3.
	// all elements in an array are automatically assigned with the zero value of the array type
	fmt.Println(a)

	b := [3]int{2,5,13} // not all elements have to be assigned: [3]int{2} returns [2 0 0]
	fmt.Println(b)

	// you can replace length with [...]. Compiler with handle it
	c := [...]int{2,1}
	fmt.Println(c)

	// the size of the array is a part of the type => arrays cannot be resized

	// arrays are value types, not reference types.
	d := [...]int{1,2}
	e := d
	e[0] = 3
	fmt.Println("d is", d)
	fmt.Println("e is", e)

	// when array are passed to function as parameters, they are passed by value
	// and the original array is unchanged

	// length - pass array to fn len()
	fmt.Println("length of e is", len(e))

	// iterate through array => normal for or range for. The latter option example:
	f := [...]int{3,43,54,32,11}
	sum := int(0)
	for i, v := range f {
		fmt.Printf("%d the element of a is %d\n", i, v)
		sum += v
	}
	fmt.Println("sum of all elements of a =", sum)

	// in case you want only value and ignore index
	for _, v := range f {
		fmt.Println("value ", v)
	}
}
