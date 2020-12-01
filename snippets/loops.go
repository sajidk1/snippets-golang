package main

import "fmt"

func main() {

	// For loop
	for i := 1; i <= 5; i++ {

		fmt.Println(i)
	}

	// Go's equivalent of a while loop
	i := 1

	for i <= 5 {

		fmt.Println(i)
		i++
	}

	// Nested loops
	for i := 1; i < 5; i++ {

		for j := 1; j < i; j++ {

			fmt.Println(j)
		}
	}

}
