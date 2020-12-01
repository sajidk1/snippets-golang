package main

import "fmt"

func main() {

	x := 5

	fmt.Println(x)

	// Get the pointer i.e. the memory address of x
	fmt.Println(&x)

	// We can change the value of x using another function using pointers. The function will change the value at the memory address
	changeValue(&x)
	fmt.Println(x)
}

func changeValue(x *int) {

	*x = 7
}
