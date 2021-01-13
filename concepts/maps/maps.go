package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":   "Nat",
		"course": "System's Analysis and Development",
	}

	fmt.Println(user)
	delete(user, "name")
	fmt.Println(user)

	// para acessar user["name"]
}
