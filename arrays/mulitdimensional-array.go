package main

import "fmt"

func printarray(a [3][2]string)  {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printarray(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung" // the rest of the elements will be field with "" (not null)
	printarray(b)
}
