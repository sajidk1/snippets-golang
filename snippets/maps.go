package main

import "fmt"

func main() {

	// Simple map
	StudentAge := make(map[string]int)

	StudentAge["Alice"] = 21
	StudentAge["Bob"] = 19
	StudentAge["Carol"] = 20
	StudentAge["Dan"] = 21
	StudentAge["Eve"] = 22

	fmt.Println(StudentAge["Alice"])
	fmt.Println(len(StudentAge))

	// Nested maps
	superhero := map[string]map[string]string{

		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},

		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham",
		},
	}

	hero := superhero["Batman"]
	fmt.Println(hero["RealName"], hero["City"])

}
