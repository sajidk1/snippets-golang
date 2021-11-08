package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	/*
		Go routine = engine that executes code i.e. a thread,
		every program starts with 1 Go routine i.e. the main control flow

		This runs concurrently, not in parallel,
		since the Go scheduler only runs one Go routine at a time.

		So what's the point then?
		When one of those routines has a blocking call/finishes, it will switch to another Go routine,
		which is still more efficient then running everything sequentially.

		If you want true parallelism, you need to use multiple CPU cores, Go by default only uses one.
	*/

	linksToCheck := []string{
		"https://stackoverflow.com",
		"https://google.com",
		"https://golang.org",
		"https://reddit.com",
		"https://github.com",
		"https://udemy.com",
		"https://slack.com",
	}

	/*
		Channels are used to communicate data between Go routines,
		you must specify the type e.g. string for the data that is passed around
	*/
	c := make(chan string)

	for _, link := range linksToCheck {
		// go keyword = run this function inside a new child Go routine/thread
		go checkLink(link, c)

	}

	// Infinite loop, waits for channel message (c), which it then assigns to var l, and then runs the loop
	for l := range c {
		/*
			Go's function literal = Anonymous function in JavaScript, or Lambda in Python
			i.e. a unamed function, useful for a small bit of code
		*/
		go func(link string) {
			// Better to pause here rather than the main Go routine, otherwise you block everything
			time.Sleep(5 * time.Second)
			// This is a blocking call (within this Go routine), so it won't move on until we get a message in the channel
			checkLink(link, c)
		}(l) // Better to pass in the actual value, rather then rely on a pointer to l, since the value of l might change within the main Go routine
	}

	/*
		Alternative channel loop syntax:

		for {
			// This is a blocking call, so it won't move on until we get a message in the channel
			go checkLink(<-c, c)
		}
	*/
}

/*
	Sending messages through channels is like sending a SMS
	channel <- 5 = send the value '5' into the channel
	myIntVar <- channel = receive value from channel
	fmt.Println(<- channel) = wait for a value from channel and print as soon as we get it

	FYI: receiving messages from a channel is a blocking call
*/
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Once we've checked the site, add it to the channel so it can be checked again
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// Once we've checked the site, add it to the channel so it can be checked again
	c <- link
}
