package main

import "fmt"

type bot interface {
	/*
		This associates any type that has a getGreeting() function which returns a string
		with type 'bot'.
		That means that type can be used with any function which expects the type 'bot'.
	*/
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {

	/*
		Interfaces help you re-use functions with different types
	*/

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	// Imagine this had custom logic specific to englishBot i.e. we don't want to reuse it for the other bot
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	// Imagine this had custom logic specific to spanishBot i.e. we don't want to reuse it for the other bot
	return "Hola"
}

func printGreeting(b bot) {
	// This has no custom logic so we don't want to repeat this code
	fmt.Println(b.getGreeting())
}
