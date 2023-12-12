package main

import (
	"command-line-application/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
