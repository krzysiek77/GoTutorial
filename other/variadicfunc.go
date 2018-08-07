package main

import "fmt"

// variadic func - func that can accept variable (any: 0, 1, 2, ...) number of arguments:
//func append(slice []Type, elems ...Type) []Type

// The way variadic functions work is by converting the variable number of arguments passed,
// to a new slice of the type of the variadic parameter.

// Is there a way to pass a slice to a variadic func?
// add (...) as a suffix of the slice.
// The slice is then passed as an argument without creating new slice.
// It means that any changes made to this slice inside of func
// will be visible outside of it as well:
//	nums := []int{85,55,67}
//	find(34, nums...)

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground") // here new slice is created, which means the change would not be visible outside the func
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
