package main

import "fmt"

// tells other types there is a new type in program
type bot interface {
	// if a type has a function called "getGreeting()" that returns a string,
	// it is now an honorary type of "bot" as well as the original type
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//pretend getGreeting has very custom/complex logic for English-specific greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	//custom logic
	return "Hola!"
}
