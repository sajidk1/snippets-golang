package main

import "fmt"

func main() {

	// https://blog.golang.org/defer-panic-and-recover

	// Defer is like "finally" in other languages, it will run after other functions return/fail
	defer FirstRun()
	SecondRun()

	// Recover is like catch. Divide by 0 to cause an error
	fmt.Println(div(3, 0))

	// Panic is like throw
	demoPanic()
}

func FirstRun() {
	fmt.Println("First")
}

func SecondRun() {
	fmt.Println("Second")
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2

	return solution

}

func demoPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}
