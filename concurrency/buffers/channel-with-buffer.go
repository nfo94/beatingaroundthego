package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Hello"

	msg := <-channel
	fmt.Println(msg)
}
