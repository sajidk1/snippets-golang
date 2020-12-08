package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// https://www.youtube.com/watch?v=LvgVSSpwND8
	// Concurrency is about breaking up a program into concurrent steps that can be parallelised (i.e. run at the same time)

	// A WaitGroup waits for a collection of goroutines to finish https://golang.org/pkg/sync/#WaitGroup
	// AKA "blocking" in other languages: https://nodejs.org/en/docs/guides/blocking-vs-non-blocking/#blocking
	var wg sync.WaitGroup
	wg.Add(1)

	// A goroutine is a lightweight thread managed by the Go runtime. https://tour.golang.org/concurrency/1
	// Anonymous function (i.e. inline function) https://www.geeksforgeeks.org/anonymous-function-in-go-language/
	go func() {
		count("sheep")
		wg.Done()
	}()

	// This will block until the Go routine is finished
	wg.Wait()

}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}
