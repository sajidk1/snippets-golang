package main

import "fmt"

func main() {

	// A constant (or literal) is a variable with a fixed value i.e. cannot be modified after its definition
	const pi float64 = 3.14159265359

	fmt.Println(pi)

	// You can declare vars like this
	var (
		varA = 1
		varB = 2
	)

	fmt.Println(varA, varB)

	// Or let Go figure it out the type by using :=
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

	varString := "string"
	varBool := true
	varInt := 1

	fmt.Println(varString, varBool, varInt)

	// Assign values to multiple vars
	x, y := 14, 15

	fmt.Println(x, ",", y)

	// Get string length
	var message string = "Hello, World!"

	fmt.Println(len(message))

	// Combine strings
	messagePart1 := "Hello,"
	messagePart2 := " World!"

	fmt.Println(messagePart1 + messagePart2)
}
