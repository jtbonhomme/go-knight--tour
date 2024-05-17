package main

import (
	"log"

	"github.com/jtbonhomme/go-knight-tour/internal/game"
)

func main() {
	var err error
	g := game.New()

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Exit")
}
