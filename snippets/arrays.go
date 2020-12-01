package main

import "fmt"

func main() {

	// Arrays = collection of vars of the same type
	var EvenNumbers [5]int

	EvenNumbers[0] = 0
	EvenNumbers[1] = 2
	EvenNumbers[2] = 4
	EvenNumbers[3] = 6
	EvenNumbers[4] = 8

	fmt.Println(EvenNumbers[2])

	// Another way to initialise arrays
	oddNumbers := [5]int{1, 3, 5, 7, 9}

	fmt.Println(oddNumbers[2])

	// Iterate over an array
	for _, value := range oddNumbers {
		fmt.Println(value)
	}

	// Array values and the index of the loop
	for i, value := range oddNumbers {
		fmt.Println(value, i)
	}

	// Slicing arrays
	numSlice := []int{0, 1, 2, 3, 4}

	sliced := numSlice[1:3]
	fmt.Println(sliced)

	// Or This assume 0 as the low index
	sliced = numSlice[:3]
	fmt.Println(sliced)

	// Arrays with make: https://tour.golang.org/moretypes/13

	// Will create a zeroed array with a capacity of 10
	makeArray := make([]int, 5, 10)

	// Copy values from another array
	copy(makeArray, numSlice)

	fmt.Println(makeArray)

	// Append values

	appendArray := append(numSlice, 3, 0, -1)

	fmt.Println(appendArray)

}
