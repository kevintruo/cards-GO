package main

import "fmt"

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
