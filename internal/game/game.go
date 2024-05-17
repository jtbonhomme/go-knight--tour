package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game manages all internal game mechanisms.
type Game struct {
	ScreenWidth     int
	ScreenHeight    int
	BackgroundColor color.Color
}

// New creates a new game object.
func New() *Game {
	g := &Game{
		ScreenWidth:     500,
		ScreenHeight:    500,
		BackgroundColor: color.RGBA{0x0b, 0x0d, 0x00, 0xff},
	}

	return g
}

// Run game loop.
func (g *Game) Run() error {

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Knight Tour")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return ebiten.RunGame(g)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
