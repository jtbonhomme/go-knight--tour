package game

import (
	"errors"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var ErrQuit = errors.New("QUIT")

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.updateBlink()

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		log.Printf("key press ESCAPE: exit program")
		return ErrQuit
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		switch g.state {
		case Started:
			g.state = Running
			g.start = time.Now()
			g.runResult = g.Knight.Run()
			log.Printf("game: run")
		}
	}

	if g.state == Running {
		select {
		case result := <-g.runResult:
			log.Println("received message from knight solver", result)
			if result {
				g.state = GameWon
			} else {
				g.state = GameLost
			}
		default:
		}
	}

	return nil
}

func (g *Game) updateBlink() {
	if g.blinkFrameCounter > BlinkFrameRate {
		g.blink = !g.blink
		g.blinkFrameCounter = 0
	}
	g.blinkFrameCounter++
}
