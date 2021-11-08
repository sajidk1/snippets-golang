package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	deckLength := 16
	if len(deck) != deckLength {
		// https://gobyexample.com/string-formatting
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(deck))
	}

	firstCard := "Ace of Spades"
	if deck[0] != firstCard {
		t.Errorf("Expected first card to be %v, but got %v", firstCard, deck[0])
	}

	lastCard := "Four of Clubs"
	// last element in the slice
	if deck[len(deck)-1] != lastCard {
		t.Errorf("Expected last card to be %v, but got %v", lastCard, deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// Attempt cleanup just incase previous tests failed and never got to cleanup stage
	os.Remove("_decktesting")

	// Create a new deck and save to file
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// Now create a deck from that saved file
	loadedDeck := newDeckFromFile("_decktesting")

	deckLength := 16
	if len(loadedDeck) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(loadedDeck))
	}

	// Cleanup
	os.Remove("_decktesting")

}
