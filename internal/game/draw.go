package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	BoardX        float32 = 50
	BoardY        float32 = 50
	BoardCellSize float32 = 50
	BoardWidth    float32 = BoardCellSize * 8
	BoardHeight   float32 = BoardCellSize * 8
)

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)

	g.drawFrame(screen)
}

func (g *Game) drawFrame(screen *ebiten.Image) {
	var x, y, width, height, strokeWidth, step float32

	x, y = BoardX, BoardY
	width, height = BoardWidth, BoardHeight
	strokeWidth = 1
	step = BoardCellSize

	for i := 0; i < 9; i++ {
		vector.StrokeLine(screen,
			x, y+step*float32(i), x+width, y+step*float32(i),
			strokeWidth, color.RGBA{0x8b, 0x8d, 0x80, 0xff}, false)
		vector.StrokeLine(screen,
			x+step*float32(i), y, x+step*float32(i), y+height,
			strokeWidth, color.RGBA{0x8b, 0x8d, 0x80, 0xff}, false)
	}
}
