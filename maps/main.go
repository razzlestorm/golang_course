package main

import "fmt"

func main() {
	// keys are string, and values are string

	// vars colors map[string]string AND
	//colors := make(map[string]string) are the same

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	fmt.Println(colors)

	delete(colors, "blue")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
