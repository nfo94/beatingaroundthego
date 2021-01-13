package main

import (
	"fmt"
)

// Pessoa type
type Pessoa struct {
	name   string
	age    uint
	gender string
}

// Student struct
type Student struct {
	Pessoa
	course     string
	university string
}

func main() {
	person1 := Pessoa{"Nat", 26, "Female"}
	fmt.Println(person1)

	student1 := Student{person1, "ADS", "Unifacs"}
	fmt.Println(student1)
}
