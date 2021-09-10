package main

import "fmt"

// Create a new Type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardVals := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, val := range cardVals {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// func (receiver with a name d and type deck) funcname(args) {}
// d is like self/this
func (d deck) print() {
	for ii, card := range d {
		fmt.Println(ii, card)
	}
}
