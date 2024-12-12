package main

import (
	"fmt"
	"time"
)

func main() {
	/*done := make(chan bool) // can use single channel multiple times
	go greet("nice to meet you", done)
	go greet("How are you", done)
	//done := make(chan bool) // it creates a channel for communication to goroutine
	go slowGreet("How... are... you...", done)
	go greet("good", done)

	<- done // it wait for the chan msg, only after this execution continues of this block
			// waits till some data to came out of the channel
	<- done
	<- done
	<- done //it waits until the last data to came out and execute onwards*/
	//dones := make([]chan bool, 4) // can create a list of channels

	done := make(chan bool)
	//dones[0] = make(chan bool)
	go greet("nice to meet you", done)
	//dones[1] = make(chan bool)
	go greet("How are you", done)
	//dones[2] = make(chan bool)
	go slowGreet("How... are... you...", done)
	//dones[3] = make(chan bool)
	go greet("good", done)

	/*for _, done := range dones {
		<-done
	} // it waits until all data to came out and execute onwards*/

	/*for doneChan := range done{
		fmt.Println(doneChan)
	}*/

	for range done {
		//
	}
}

func greet(msg string, doneChan chan bool) {
	fmt.Println(" Hello,", msg)
	doneChan <- true
}

func slowGreet(msg string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(" Hello,", msg)
	doneChan <- true // sends data to a channel
	close(doneChan) // it close the channel if last data sent through the channel
}
