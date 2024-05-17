package game

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var ErrQuit = errors.New("QUIT")

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		log.Printf("key press ESCAPE: exit program")
		return ErrQuit
	}

	return nil
}
