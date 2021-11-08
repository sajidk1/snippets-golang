package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		The task is simple:
		Print integers 1 to N, but print “Fizz” if an integer is divisible by 3,
		“Buzz” if an integer is divisible by 5,
		and “FizzBuzz” if an integer is divisible by both 3 and 5.
		https://www.youtube.com/watch?v=QPZ0pIK_wsc
	*/

	// Method 1
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// Method 2
	for i := 1; i <= 100; i++ {

		var output string

		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			output = strconv.Itoa(i)
		}

		fmt.Println(output)
	}

}
