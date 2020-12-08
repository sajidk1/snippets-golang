package main

import (
	"fmt"
	"time"
)

func main() {

	// More here: https://github.com/jakewright/tutorials/tree/master/go/02-go-concurrency

	// Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine and receive those values into another goroutine.
	// https://tour.golang.org/concurrency/2
	c := make(chan string)
	go count("sheep", c)

	// Receiving/Sending to a channel is a blocking operation so it allows you to synchronies go routines
	// https://youtu.be/LvgVSSpwND8?t=462
	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	// As a sender: once done close the channel to ensure receivers are not waiting for messages (avoid deadlock)
	// https://gist.github.com/YumaInaura/8d52e73dac7dc361745bf568c3c4ba37#deadlock
	close(c)

}
