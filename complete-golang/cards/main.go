package main

//import "fmt"

func main() {

	/* 	cards := newDeck()

	   	// Assigning multiple returned values to two variables
	   	hand, remainingCards := deal(cards, 5)

	   	hand.print()
	   	remainingCards.print()

	   	cards.saveToFile("cards.txt") */
	cards := newDeckFromFile("cards.txt")

	cards.print()

	cards.shuffle()

	cards.print()

}
