package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// don't necessarily have to write out field name
	contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@test.com",
			zipCode: 98025,
		},
	}

	// Go is a pass-by-value language
	// The following line does not update the alex person's name: due to Go creating a copy in another mem address when a function is called
	// alex.updateName("Alexis")
	// &var = give me memory address of the value this variable is pointing at.
	// alex = value associated with alex, &alex = memory address with alex
	alexPointer := &alex
	alexPointer.updateName("Alexis")
	// %+v prints out all field names and their values
	alex.print()
	alex.updateName("AlexisNoPointer")
	alex.print()
}

// *pointer = give me the value in the memory address this is pointing at
// whenever you see a * where a type should be, it is type[]pointerToType

// This function can also be called with both a pointer to person or a person type.
func (pointerToPerson *person) updateName(newFirstName string) {
	// Below is an operator that takes this pointer and turns it into an actual value.
	(*pointerToPerson).firstName = newFirstName
}

// You can turn Address into Value with *Address
// You can turn Value into Address with &Value
// Slices don't need to worry about pointers, it's mainly structs?

func (p person) print() {
	fmt.Printf("%+v", p)
}

// When creating a slice, Go creates two separate data structures.
// The first is a Slice (with pointer to Array head, capacity, and length), and then an internal array that represents all of the items.
// use pointers to change the following in a function: int, float, string, bool, struct (these are value types)
// pointers not necessary for: slices, maps, channels, pointers, functions (these are reference types)
