package main

import (
	"log"
	"os"

	"github.com/fajarihsan21/go-backend/src/configs/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
