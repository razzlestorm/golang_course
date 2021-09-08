// package == project == workspace
// Go has two packages: executable (go build) and reusable (code libraries)
// package main == making executable package, anything else is for a reusable package
package main

import "fmt"

// package main must have a main func
func main() {
	fmt.Println("Hello world!")
}
