package main

import (
	"log"

	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenReSTTree(test, "./")
	if err != nil {
		log.Fatal(err)
	}
}
