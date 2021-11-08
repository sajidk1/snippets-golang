package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Takes first arg e.g. 'go run main.go test.txt'
	fileName := os.Args[1]

	// Returns a pointer to a file or an error if there is one
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// io.Copy(thewriter, thereader)
	io.Copy(os.Stdout, file)

}
