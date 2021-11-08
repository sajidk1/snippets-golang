package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

// You need to declare the type you are going to return
func newDeck() deck {

	cards := deck{}

	// Array = Fixed length list of things
	// Slice = Array that can grow or shrink
	// Arrays/Slices can only contain elements of the same type
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Use _, if you don't need the index variable within your loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// Append does not modify the existing slice, rather it returns a value,
			// so you should assign it to your slice variable
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function, allows to you add methods to custom types
// Convention is to use 1 or 2 letter abbreviation for the receiver (d), that matches the type name (deck)
func (d deck) print() {

	// i = index of current item, card = current item in loop
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Function signature: you must specify types for paramater and returned values
func deal(d deck, handSize int) (deck, deck) {

	// multiple return values of type deck
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {

	// Type conversion to slice of strings
	slice := []string(d)

	// Then convert the slice to one long comma delimited string
	return strings.Join(slice, ",")

}

func (d deck) saveToFile(filename string) error {
	// Convert deck into a string
	string := d.toString()

	// Convert string to byte slice (string as UTF-8 encodings/chracters)
	byteSlice := []byte(string)

	// Write file to disk, giving everyone read/write permission
	return ioutil.WriteFile(filename, byteSlice, 0666)
}

func newDeckFromFile(filename string) deck {

	// Read file from disk
	byteSlice, error := ioutil.ReadFile(filename)

	// Check for errors
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	// Otherwise....

	// Convert byte slice to a string
	string := string(byteSlice)

	// Convert string to slice of strings
	sliceOfStrings := strings.Split(string, ",")

	// Convert slice of strings to deck and return
	return deck(sliceOfStrings)

}

func (d deck) shuffle() {

	// Generate a seed
	now := time.Now().UnixNano()
	source := rand.NewSource(now)
	r := rand.New(source)

	for i := range d {

		// Generate a random number
		newPosition := r.Intn(len(d) - 1)

		// Swap the index of the current card with the random number
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
