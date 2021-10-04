package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	// Declare deck
	cards := deck{}

	// Declare cardSuits and cardValues to iterate and create a card and append to cards each time a card is created -> avoid tedious work
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Create card and append to cards
	// NEW: replace indices (ie. i, j) with _ to tell the GO compiler that we are not going to use iterate
	// GO will throw an error if a variable it not used
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Clarity to the method deal() below.
// For example: fruits := []string{"apple", "orange", "banana", "grape"}
// fruits[0] = "apple", fruits[1] = "orange", etc.
// There are two ways to use slice range, they are both equivalent, most people prefer the second method
// The number 2 here means that it is up to index number 2 but not including the value at index number 2
// For example:
// fruits[0:2] = {"apple", "orange"} or fruits[:2] = {"apple", "orange"}
// We can also use it the other way
// For example:
// fruits[2:] = {"banana", "grape"}
// In our case, the handSize is the number '2' in the example!
// NEW: you can return multiple values in GO!
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString method just like Java, but you have to do it manually
// Learned about helper function Join, basically adding a delimeter to the slice
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
