package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrencia != paralelismo
	// to run both you can use goroutine
	go write("Hello")
	write("Programming")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
