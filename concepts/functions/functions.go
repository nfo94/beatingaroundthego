package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n1
}

func maths(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	somar := sum(10, 20)
	fmt.Println(somar)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fmt.Println(f("Testando"))

	fmt.Println(maths(20, 15))
}
