package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting1(eb)
	printGreeting2(sb)

	/*
		Life would be easier if we could use a generic printGreeting for
		both types! Interfaces to the rescue.
	*/

}

func (eb englishBot) getGreeting() string {
	// Imagine this had custom logic specific to englishBot i.e. we don't want to reuse it for the other bot
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	// Imagine this had custom logic specific to spanishBot i.e. we don't want to reuse it for the other bot
	return "Hola"
}

func printGreeting1(eb englishBot) {
	// This has no custom logic and ideally we don't want to repeat this code
	fmt.Println(eb.getGreeting())
}

func printGreeting2(sb spanishBot) {
	// This has no custom logic and ideally we don't want to repeat this code
	fmt.Println(sb.getGreeting())
}
