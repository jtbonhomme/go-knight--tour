package main

import (
	"flag"
	"log"

	"github.com/jtbonhomme/go-knight-tour/internal/game"
)

func main() {
	var err error
	var speed int

	flag.IntVar(&speed, "s", 1, "speed resolution (default is 1)")
	flag.Parse()

	g := game.New(speed)

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Exit")
}
