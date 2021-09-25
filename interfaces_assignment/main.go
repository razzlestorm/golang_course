package main

import "fmt"

// Write a program that creates two custom struct types called triangle and square

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		height: 1.5,
		base:   1.0,
	}
	s := square{sideLength: 4.3}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(fmt.Sprintf("shape area is %f.", s.getArea()))
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return (s.sideLength * s.sideLength)
}
