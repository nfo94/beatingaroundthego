package main

import (
	"first-module/auxiliary"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello world")
	auxiliary.Write()
	error := checkmail.ValidateFormat("nat@gmail.com")
	fmt.Println(error)
}
