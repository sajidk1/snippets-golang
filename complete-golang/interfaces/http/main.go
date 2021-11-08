package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp.Body) // This won't work

	/*
		Fields in a struct can be an interface: e.g. 'Body io.ReadCloser' (https://pkg.go.dev/net/http@go1.17.2#Response)
		Interfaces can reference other interfaces: e.g
		type ReadCloser interface {
			Reader
			Closer
		}
		(https://pkg.go.dev/io#ReadCloser)
	*/

	/*
		// Working with the Reader yourself:
		byteSlice := make([]byte, 99999) // number of elements/empty spaces as Read does not auto resize the byte slice
		resp.Body.Read(byteSlice)
		fmt.Println(string(byteSlice))
	*/

	// Using helper functions which work with the Writer interface
	io.Copy(os.Stdout, resp.Body)

	// A custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	byteSliceLength := len(bs)
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", byteSliceLength)
	return byteSliceLength, nil
}
