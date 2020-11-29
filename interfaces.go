package main

import (
	"fmt"
	"math"
)

func main() {

	// Interfaces explained: https://stackoverflow.com/a/6878726
	// TL;DR A method that has a consistent "interface" (syntax, arguments) but allows you to have different implementations for different types

	// https://gobyexample.com/interfaces

	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of rectangle is", getArea(rect))

	fmt.Println("Area of circle is", getArea(circ))

}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {

	return r1.height * r1.width
}

func (c1 Circle) area() float64 {

	return math.Pi * math.Pow(c1.radius, 2)
}

func getArea(shape Shape) float64 {

	return shape.area()
}
