package main

import "fmt"

// if consecutive parameters are of the same type, we can avoid writing the typ each time
// and it is enough to be written once at the end.
// func calculateBill(price, no int) int {...}
func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// func can retrun multiple values
func rectProps(length, width float64)(float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func main() {
	calculateBill(10, 5)

	a, b := rectProps2(10, 20)
	fmt.Println("area: ", a, "perimeter: ", b)
}

// func can return named values. It's considered than as declaring variable(s) in the 1st lie of the functions
func rectProps2(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no explicit return value
}

// _ is the blank identifier and can be used in place of any value of any type
// e.g.: to discard not needed value (rectProps2 - we want area, but dont give a fuck about perimeter)

