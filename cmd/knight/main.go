package main

import (
	"flag"
	"log"

	"github.com/jtbonhomme/go-knight-tour/internal/game"
)

func main() {
	var err error
	var slowMotion int
	var implementation string
	var debug bool

	flag.StringVar(&implementation, "i", "naive", "implementation (default is \"naive\")")
	flag.IntVar(&slowMotion, "s", 1, "slowMotion resolution (default is 1)")
	flag.BoolVar(&debug, "d", false, "debug (default is false)")
	flag.Parse()

	if slowMotion < 1 {
		slowMotion = 1
	}
	g := game.New(slowMotion, implementation, debug)

	log.Println("Start game")
	err = g.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Exit")
}
