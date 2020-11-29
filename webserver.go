package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/page1", Page1Handler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func Page1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page 1\n")
}
