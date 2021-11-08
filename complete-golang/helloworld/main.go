/*
package = project or workspace, a collection of common source code files
every .go files need to declare what package it belongs to
if executable (i.e. you want to use 'go build' it's 'package main',
otherwise it's a reusable package (i.e. helper/library) then 'package mymodule'
*/
package main

/*
Go has many standard packages: https://pkg.go.dev/std
You can of course import other local or external packages
*/
import "fmt"

/*
func = function
Just like any other language
Use the keyword to declare a function (func)
Set the function name e.g. main, getStuff
In the parentheses list the arguments you can pass to the function
*/
func main() {
	fmt.Println("Hello, World!")
}
