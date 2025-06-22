package main

import (
	"log"
	"os"

	"github.com/gatsu420/git-email-collector/app"
)

func main() {
	err := app.Collect(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
