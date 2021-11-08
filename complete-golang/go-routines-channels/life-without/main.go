package main

import (
	"fmt"
	"net/http"
)

func main() {

	// This runs sequentially

	linksToCheck := []string{
		"https://stackoverflow.com",
		"https://google.com",
		"https://golang.org",
		"https://reddit.com",
		"https://github.com",
		"https://udemy.com",
		"https://slack.com",
	}

	for _, link := range linksToCheck {
		checkLink(link)
	}

}

func checkLink(link string) {
	/*
		On each iteration of the loop we are waiting
		for 'http.Get' to complete before moving on,
		i.e. it's a blocking call, that's inefficient
	*/
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
