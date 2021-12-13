package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	height float64
}
type square struct {
	sideLength float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}