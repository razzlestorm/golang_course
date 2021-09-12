package main

// Notes:
// Array in Go is fixed length
// Slice can grow/shrink in size
// Both must be defined by a data type

func main() {
	// deck defined in deck.go
	cards := newDeck()
	// append not modifying the original object is an important point in Go

	cards.saveToFile("my_cards")
}

// Write datatype being returned just after arguments parens
