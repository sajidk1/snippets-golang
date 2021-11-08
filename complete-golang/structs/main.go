package main

import "fmt"

// Struct = data structure, collection of properties that are realted together
type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	/*

		// This relies on you providing the values in the correct order
		alex := person{"Alex", "Anderson"}

		// This doesn't
		alex := person{lastName: "Anderson", firstName: "Alex"}
	*/

	// Variables declared without an explicit initial value are given their respective zero value.
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	alex.print()

	// One way of declaring values for embedded structs
	jim := person{
		firstName: "Jim",
		lastName:  "Johnson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

// Receiver functions

func (p person) print() {
	// Print field names aswell as values
	fmt.Printf("%+v", p)
}

/*
Go is a "pass by value" language,
meaning that a copy of the actual parameter's value is made in memory,
i.e. the caller and callee have two independent variables with the same value.
More info here: https://stackoverflow.com/questions/373419/whats-the-difference-between-passing-by-reference-vs-passing-by-value
To update the caller's value you must use pointers
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
