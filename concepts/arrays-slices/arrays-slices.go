package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello")

	var arr1 [3]string
	arr1[0] = "First"
	arr1[1] = "Second"
	arr1[2] = "Third"
	fmt.Println(arr1)

	arr2 := [2]string{"First", "Second"}
	fmt.Println(arr2)

	slice := []int{2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := arr2[1:2]
	fmt.Println(slice2)
	fmt.Println(slice[0])
}
