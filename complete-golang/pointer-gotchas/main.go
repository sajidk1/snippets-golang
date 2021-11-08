package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

/*
Although Go is still making a copy of the slice when passed in,
the slice type itself only has a pointer to the actual data,
and does not hold the data itself.
That's why you can update a slice without a pointer!
*/
func updateSlice(s []string) {
	s[0] = "Bye"
}

/*
Types that behave like this are called 'Reference' types,
they include: slices, maps, channels, pointers, functions

Otherwise, it's a 'Value' type (where you need to worry about pointers),
they include: int, float, string, bool, structs
*/
