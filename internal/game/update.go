package game

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var ErrQuit = errors.New("QUIT")

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		log.Printf("key press ESCAPE: exit program")
		return ErrQuit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		switch g.state {
		case Started:
			g.state = Running
			g.Knight.Run()
			log.Printf("game: run")
		case Running:
			g.state = Paused
			g.Knight.Pause()
			log.Printf("game: pause")
		case Paused:
			g.state = Running
			g.Knight.Run()
			log.Printf("game: run")
		}
	}
	return nil
}
