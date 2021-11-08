package main

import "fmt"

func main() {

	/*
		Map = collection of key/value pairs
		All keys must be of the same type, and all values must be of the same type
		Key and value can be of different types
	*/

	// Declaring maps
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#33cc33",
	}
	fmt.Println(colours)

	var colours2 map[string]string
	fmt.Println(colours2)

	colours3 := make(map[string]string)
	fmt.Println(colours3)

	// Accessing/adding keys, you must use square brackets
	redHexCode := colours["red"]
	fmt.Println(redHexCode)
	colours["white"] = "#ffffff"
	fmt.Println(colours)

	printMap(colours)

	// Deleting key/value pairs
	delete(colours, "green")
	fmt.Println(colours)

}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}

/*
Use a map when you have a collection of closely related properties
(like the colours map in this file),
and all your keys are not known at compile time.

Otherwise, use a struct
*/
