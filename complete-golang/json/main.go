package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type outJsonMessage struct {
	Name string
	Body string
	Time int64
}

// https://golangbyexample.com/exported-unexported-fields-struct-go/
type inJsonMessage struct {
	UserId    int
	Id        int // JSON tags: https://golang.cafe/blog/golang-json-marshal-example.html
	Title     string
	Completed bool
}

func main() {

	// https://go.dev/blog/json and https://tutorialedge.net/golang/parsing-json-with-golang/

	// Write to .JSON file
	outJson := outJsonMessage{"Alice", "Hello", 1294706395881547000}

	outRegular, err := json.Marshal(outJson)
	if err != nil {
		handleError(err)
	}

	outPretty, err := json.MarshalIndent(outJson, "", "  ")
	if err != nil {
		handleError(err)
	}

	fmt.Println("outRegular:", string(outRegular))
	fmt.Println("outPretty:", string(outPretty))

	ioutil.WriteFile("out.JSON", outPretty, 0666)

	// Read from .JSON file
	inJson, err := ioutil.ReadFile("in.JSON")
	if err != nil {
		handleError(err)
	}

	var inJsonAsStruct inJsonMessage
	errs := json.Unmarshal(inJson, &inJsonAsStruct)
	if errs != nil {
		handleError(errs)
	}

	fmt.Println("inJsonAsStruct.Title", inJsonAsStruct.Title)
	fmt.Printf("inJsonAsStruct: %+v\n", inJsonAsStruct)

	// Read from .JSON without knowing the structure beforehand
	var unknownJson map[string]interface{}

	errors := json.Unmarshal(inJson, &unknownJson)
	if errors != nil {
		handleError(errors)
	}

	// https://irshadhasmat.medium.com/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
	fmt.Println("unknownJson[title]:", unknownJson["title"])
	fmt.Println("unknownJson:", unknownJson)

	// JSON from API
	// https://tutorialedge.net/golang/consuming-restful-api-with-go/

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		handleError(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		handleError(err)
	}
	fmt.Println("responseData:", string(responseData))

}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
