package main

import (
	"fmt"
	"go-cli/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	application := app.Generate()

	application.Run(os.Args)
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
