package main

import "fmt"

func main() {
	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	x := 42.0

	switch x {
	case 0, 1:
		fmt.Println("Zero or one")
	case 42:
		fmt.Println("42")
	default:
		fmt.Println("default")
	}

	for key, value := range map[string]int{"one": 2, "two": 2} {
		fmt.Printf("key=%s, value=%d/n", key, value)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("iteration", i)
	}
}
