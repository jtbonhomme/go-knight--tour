package main

import (
	"flag"
	"log"

	"github.com/jtbonhomme/go-knight-tour/internal/game"
)

func main() {
	var err error
	var speed int
	var implementation string

	flag.StringVar(&implementation, "i", "naive", "implementation (default is \"naive\")")
	flag.IntVar(&speed, "s", 0, "speed resolution (default is 0)")
	flag.Parse()

	g := game.New(speed, implementation)

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Exit")
}
