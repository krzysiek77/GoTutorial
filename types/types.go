package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// int32 or int64
	var a int = 89
	b := 95
	fmt.Printf("type of a is %T, size of a is %d.\n", a, unsafe.Sizeof(a))
	fmt.Printf("type of b is %T, size of b is %d.\n", b, unsafe.Sizeof(b))

	// floats
	c, d := 5.67, 8.97
	fmt.Printf("type of c is %T, d is %T\n", c, d)

	// complex64 (float32 + imaginary part)
	// complex128 (float64 + imaginary part)
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("product: ", cmul)

	// there is no automatic type promotion or conversion
	i := 55 //int
	j := 67.8 //float64
	//sum : i + j //not allowed
	sum := i + int(j)
	fmt.Println(sum)

	x := 54
	var y float64 = float64(x)
	fmt.Println("y = ", y)
}
