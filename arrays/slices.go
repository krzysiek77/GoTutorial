package main

import "fmt"

// slices - flexible wrapper on top of an array
//	don't own any data
//	just a references to existing arrays

func main() {
	a := [5]int{1,3,34,21,45}
	var b []int = a[1:4] // slice from a[1] to a[3] (a[start:end] = a[start]..a[end-1])
	fmt.Println(b)

	// other way to create slice
	c := []int{6,7,8}
	fmt.Println(c)

	// change to slice is a change to underlying array
	// if more than one slice share the same underlying array,
	//	the changes to each one will be reflected in the array (and among slices)
	nums := [3]int{1,2,3}
	nums1 := nums[:] // slice of all array elements
	nums2 := nums[:]
	fmt.Println("before any changes: array", nums, "slice 1", nums1, "slice 2", nums2)
	nums1[0] = 11
	fmt.Println("after changes to slice 1: array", nums, "slice 1", nums1, "slice 2", nums2)
	nums2[1] = 22
	fmt.Println("after changes to slice 2: array", nums, "slice 1", nums1, "slice 2", nums2)

	// length - number of elements in the slice
	// capacity - number of elements in the underlying array starting from the index from which the slice is created
	d := [...]int{1,2,3,44,53,456}
	e := d[1:4]
	fmt.Printf("length of slice %d, capacity %d\n", len(e), cap(e))

	// func make([]T, len, cap)
	//	- crates an arraye and returns a slice reference to it
	//	- cap param is optional and defaults to len
	i := make([]int, 5, 5)
	j := make([]int, 5)
	fmt.Println(i)
	fmt.Println(j)

	// func append(s []T, x ...T) []T
	//	- x ...T - variable number of argument for the parameter x (variadic function)
	//
	// what happens when appending:
	//	- new array is created
	//	- elements of old array are copied over to the new one
	//	- new slice to the new array is returned
	//	- slice capacity is doubled
	cars := []int{1,2,3}
	// (1)
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, 4)
	// (2) - capacity of (1) doubled
	fmt.Println("cars", cars, "has new length", len(cars), "and capacity", cap(cars))
	cars = append(cars, 5)
	// (3) - capacity the same as for (2), because there are more free elements in underlying array than needed
	fmt.Println("cars", cars, "has new length", len(cars), "and capacity", cap(cars))
	cars = append(cars, 6)
	cars = append(cars, 7)
	// (4) - capacity of (3) doubled
	fmt.Println("cars", cars, "has new length", len(cars), "and capacity", cap(cars))

	//  - zero vale of a slice type is nil
	//	- nil slice has length and capacity 0
	//	- it's possible to append values to a nil slice with append func.
	var names []string //zero value of a slice is nil
	fmt.Println("names", names, "length", len(names), "capacity", cap(names))

	// it's possible to append one slice to another using ... operator
	veggies := []string{"potatoes","tomatoes","brinjal"}
	fruits := []string{"oranges","apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)

	// slice can be thought of as being represented internally by struct:
	type slice struct {
		Length			int
		Capacity		int
		ZerothElement	*byte
	}

	// when a slice is passed to a function, even though it's passed by value,
	// the pointer variable will refer to the same underlying array
	nos := []int{8,7,6}
	fmt.Println("slice before fn call", nos)
	substractOne(nos)
	fmt.Println("slice after fn call", nos)

	// slices can be multidimensional
	//var a = [][]int{}

	// Slices hold a reference to the underlying array. As long as the slice is in memory,
	// the array cannot be garbage collected. This might be of concern when it comes to memory management.
	// Lets assume that we have a very large array and we are interested in processing only a small part of it.
	// Henceforth we create a slice from that array and start processing the slice.
	// The important thing to be noted here is that the array will still be in memory since the slice references it.

	// One way to solve this problem is to use the copy function:
	// 	func copy(dst, src []T) int
	// to make a copy of that slice.
	// This way we can use the new slice and the original array can be garbage collected.
}

func substractOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
