package main

import (
	"fmt"
)

// Usuario struct
type Usuario struct {
	name string
	age  uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var user1 Usuario
	user1.name = "Nat"
	user1.age = 26
	fmt.Println(user1)

	user2 := Usuario{"Davi", 21}
	fmt.Println(user2)
}
