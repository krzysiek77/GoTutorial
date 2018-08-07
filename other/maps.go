package main

import "fmt"

// map
// 	- builtin type that associates a value to a key
//	- similar to slices, maps are reference types
//	- maps can't be compared to each other with ==. It works only for comparison with 'nil'
//
// how to create a map:
//	make(map[type of key]type of value)
//
// the zero value of map is nil.
//
// adding item to nil map results in run time panic.
//	hence the map has to be initialized using 'make'

func initialize_map() map[string]int {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. initializing...")
		personSalary = make(map[string]int)
	}
	return personSalary
}

func add_new_item(salaries map[string]int, key string, val int) {
	salaries[key] = val
}


func main()  {
	salaries := initialize_map()
	fmt.Println("salaries", salaries)
	add_new_item(salaries, "John", 26000)
	fmt.Println("add John", salaries)
	add_new_item(salaries, "Megan", 33000)
	fmt.Println("add Megan", salaries)

	// initialize map during declaration:
	personSalary := map[string]int {
		"John": 30000,
		"Roman": 43000,
	}
	fmt.Println("person salary", personSalary)

	// if the element doesn't exist in map - zero value of elements type will be returned

	// how to check if key is present in a map?
	value, ok := personSalary["John"]
	fmt.Println("John is in a map", ok, " with salary", value)
	value, ok = personSalary["Ala"]
	fmt.Println("Ala is in a map", ok, " with salary", value)

	// how to iterate over the map?
	// 	!!! the order in which the values are retrieved is not guaranteed for the 'for range'
	for key, value := range personSalary {
		fmt.Println(key, value)
	}

	// how to delete element?
	//	delete(map, key)
	//	delete doesn't return any value
	delete(personSalary, "John")
	fmt.Println("Person Salary after deleting John", personSalary)
	delete(personSalary, "Kate")
	fmt.Println("Person Salary after deleting Kate (who doesn't exist there)", personSalary)

	// length of the map?
	//	len(mapName)
}


