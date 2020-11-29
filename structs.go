package main

import "fmt"

func main() {

	// A struct (short for "structure") is a collection of data fields with declared data types
	// Structs can improve modularity and allow to create and pass complex data structures around the system
	// See https://www.golangprograms.com/go-language/struct.html and https://gobyexample.com/structs

	rectangle1 := Rectangle{height: 10, width: 5}
	rectangle2 := Rectangle{10, 5}

	fmt.Println(rectangle1.height)
	fmt.Println(rectangle2.width)

	fmt.Println(rectangle1.area())
	fmt.Println("Area of Rectangle is", rectangle2.area())
}

type Rectangle struct {
	height float64
	width  float64
}

// *Rectangle = The struct the function will accept
func (rect *Rectangle) area() float64 {

	return rect.width * rect.height
}
