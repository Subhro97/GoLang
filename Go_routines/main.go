package main

import (
	"fmt"
	"time"
)

func main() {
	doneChans := make([]chan bool, 4) // Create a slices for storing channels

	for index := range doneChans {
		doneChans[index] = make(chan bool) // Replacing the nil values in the slices with channels to communicate with the go routines
	}

	// Converting all the functions into go routines
	go greet("Hello!", doneChans[0])
	go greet("How are you doing?", doneChans[1])
	go slowGreet("This is a slow Greeting!", doneChans[2])
	go greet("I am doing great!!", doneChans[3])

	// Only when all the routines return true from the channel, does the main method complete execution
	for index := range doneChans {
		<-doneChans[index]
	}
}

func greet(value string, doneChan chan bool) {
	fmt.Println(value)

	doneChan <- true
}

func slowGreet(value string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(value)

	doneChan <- true
}
