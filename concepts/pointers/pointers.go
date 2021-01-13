package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// a pointer is a reference to memory
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, *ponteiro)
}
