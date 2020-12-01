package main

import (
    "fmt"
    "rsc.io/quote"
    "github.com/sajid-khan-js/snippets-golang/modules/greetings"
)

func main() {
    
    fmt.Println(quote.Go())
    message := greetings.Hello("Sajid")
    fmt.Println(message)
}
