package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

// Write datatype being returned just after arguments parens
func newCard() string {
	return "Five of Diamonds"
}
