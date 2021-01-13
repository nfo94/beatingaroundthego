package main

import (
	"errors"
	"fmt"
)

func main() {
	var variable1 string = "variable 1"
	variable2 := "variable 2"
	fmt.Println(variable1, variable2)

	var (
		variable3 string = "three"
		variable4 string = "four"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "var5", "var6"
	fmt.Println(variable5, variable6)

	var boolean1 bool
	fmt.Println(boolean1)

	var err error = errors.New("Internal error")
	fmt.Println(err)
}
