package main

import "fmt"

func main() {

	age := 18

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	// Switch statements. More examples here: https://www.dotnetperls.com/switch-go

	grade := "D"

	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Well done")
	case "C", "D":
		fmt.Println("Pass")
	case "F":
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid grade")
	}

}
