package main

import "fmt"

func main() {

	// Arithmetic
	x, y := 5, 6

	fmt.Println(x, "+", y, "=", x+y)
	fmt.Println(x, "-", y, "=", x-y)
	fmt.Println(x, "*", y, "=", x*y)
	fmt.Println(x, "/", y, "=", x/y)
	fmt.Println(x, "mod", y, "=", x%y)

	// Logical
	isBool1 := true
	var isBool2 bool = false

	fmt.Println(isBool1 && isBool2) // false
	fmt.Println(isBool1 || isBool2) // true
	fmt.Println(!isBool1)           // false

}
