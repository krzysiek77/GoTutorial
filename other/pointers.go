package main

import "fmt"

// pointer:
//	- variable that stores the memory address of another variable
//	- declaring pointer: *T - it points to a value of type T:
//	- zero value is null
//	- & - operator to get the address of variable (&a)
//	- dereferencing a pointer - accessing the value of the variable which the pointer points to: *a

func main()  {
	b := 255
	var a *int = &b // & - operator to get the address of variable
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("address of b is", a)

	zeroValue()
	dereferencing()
	passPointerToFunc()
	modifyArr()
	modifySlice()
}

func zeroValue() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

// how to get and change value from a pointer
func dereferencing() {
	b := 255
	a := &b // or: var a *int = b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", *a)
}

// passing a pointer to a func
func change(val *int) {
	*val = 55
}

func passPointerToFunc() {
	a := 58
	b := &a
	fmt.Println("value of a before function call is",a)
	change(b)
	fmt.Println("value of a after function call is",a)

}

// Do not pass a pointer to an array as an argument to a function. Use slice instead.
func modify(arr *[3]int) {
	arr[0] = 90 // that's work too: (*arr)[0] = 90
}

func modifyArr() {
	a := [3]int{89, 90, 91}
	modify(&a)
	fmt.Println(a)
}

func modify2(sls []int) {
	sls[0] = 90
}

func modifySlice()  {
	a := [3]int{89, 90, 91}
	modify2(a[:])	// pass slices, no pointers to an array!!!
	fmt.Println(a)
}
