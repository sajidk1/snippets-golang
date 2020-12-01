package main

import "fmt"

func main() {

	const pi float64 = 3.14159265359
	x := 5
	isBool := true

	// Custom print aka printf: https://golang.org/pkg/fmt/

	// Float
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)

	// Type
	fmt.Printf("%T \n", pi)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", isBool)

	// Boolean
	fmt.Printf("%t \n", isBool)

	// Integers
	fmt.Printf("%d \n", x)

	// Binary
	fmt.Printf("%b \n", 25)

	// ASCII Code
	fmt.Printf("%c \n", 33)

	// Hex
	fmt.Printf("%x \n", 15)

	// Scientific notation, e.g. -1.234456e+78
	fmt.Printf("%e \n", pi)

}
