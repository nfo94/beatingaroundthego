package main

import (
	"fmt"
	"time"
)

func main() {
	// initialze variable with make, saving a channel
	channel := make(chan string)
	// goroutine for concurrency/paralellism with the function write
	go write("Hello world", channel)

	fmt.Println("After the write function starts")

	// here the channel is waiting to receive some value
	// and this is stored in a variable
	msg := <-channel
	// now i print the variable
	fmt.Println(msg)
}

// the function that was called in line 12
func write(text string, channel chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		// this is the line that sends a value to the channel
		channel <- text
		time.Sleep(time.Second)
	}

	// we can use close() to close connections
}
