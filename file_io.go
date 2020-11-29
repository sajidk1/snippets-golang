package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	fileName := "file.txt"

	// Create instance of file
	file, err := os.Create(fileName)

	// Log if any errors
	if err != nil {
		log.Fatal(err)
	}

	// Write text to file
	file.WriteString("Hello, World!")

	file.Close()

	// Read files
	stream, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	string1 := string(stream)

	fmt.Println(string1)
}
