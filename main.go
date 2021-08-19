package main

import (
	"github.com/Freeline95/GoCrud/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}