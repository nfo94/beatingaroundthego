package main

import "fmt"

func main() {
	piece := make([]float32, 10, 11)
	fmt.Println(piece)
	fmt.Println(len(piece))
	fmt.Println(cap(piece))

	piece = append(piece, 1)
	piece = append(piece, 2)
	fmt.Println(piece)
	fmt.Println(len(piece))
	fmt.Println(cap(piece))
}
