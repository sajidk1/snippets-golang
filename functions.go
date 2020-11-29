package main

import "fmt"

func main() {

	x, y := 5, 4
	fmt.Println(add(x, y))

	// Recursion: Calling a function inside the same function: https://www.martincartledge.io/how-to-write-a-recursive-function-in-go/

	// The factorial of a number is a number multiplied by "number minus one", then by "number minus two", and so on until its "number minus n = 1"
	num := 5
	fmt.Println(factorial(num))

	// Pass arbitrary amount of arguments to a function
	fmt.Println(GetTotal(10, 20, 30, 40, 50))

}

func add(num1, num2 int) int {
	return num1 + num2
}

func factorial(num int) int {

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func GetTotal(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}

	return sum
}
