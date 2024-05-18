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
		g.Restart()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.speed++
		g.speedChange <- g.speed
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		s := g.speed - 1
		if s < 0 {
			g.speed = 0
			g.speedChange <- g.speed
		} else {
			g.speed = s
			g.speedChange <- g.speed
		}
	}

	if g.state == Running {
		g.ticks += 1 + uint64(g.speed)
		g.duration = time.Millisecond * time.Duration(1000/ebiten.ActualTPS()) / time.Duration(g.ticks)
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
