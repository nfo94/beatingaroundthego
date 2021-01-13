package main

import "fmt"

func main() {
	fmt.Println("Add + double two numbers", func(a, b int) int {
		return (a + b) * 2
	}(10, 2))
}
