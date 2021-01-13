package main

import (
	"fmt"
)

func main() {
	soma := 1 + 2
	subtracao := 5 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDaDivisao := 20 / 2

	// you can't make operations with different types
	// relational operations are just like >, <=, etc
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
}
