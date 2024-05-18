package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jtbonhomme/go-knight-tour/internal/fonts"
	"github.com/jtbonhomme/go-knight-tour/internal/knight"
	"github.com/jtbonhomme/go-knight-tour/internal/text"
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
	g.drawKnight(screen)
	g.drawFinish(screen)
	g.drawDebug(screen)
}

func (g *Game) drawDebug(screen *ebiten.Image) {
	if g.state == Started {
		text.DrawTextAtPos(
			screen, fonts.SmallFont,
			50,
			470,
			"Press SPACE to start solving knight tour problem",
			color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
		)
		return
	}

	text.DrawTextAtPos(
		screen, fonts.SmallFont,
		50,
		470,
		"Press SPACE to restart",
		color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
	)

	text.DrawTextAtPos(
		screen, fonts.SmallFont,
		50,
		490,
		fmt.Sprintf("tour: %02d     time: %s", g.Knight.Tour(), g.duration),
		color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
	)
}

func (g *Game) drawKnight(screen *ebiten.Image) {
	if g.state == Started {
		return
	}

	lastPosition := knight.Position{X: -1, Y: -1}

	for _, p := range g.Knight.Positions {
		g.drawKnightPosition(screen, p)
		if lastPosition.X != -1 && lastPosition.Y != -1 {
			g.drawKnightMove(screen, lastPosition, p)
		}
		lastPosition = p
	}
}

func getCoordinatesFromPosition(p knight.Position) (float32, float32) {
	var x, y float32
	x = BoardX + float32(p.X)*BoardCellSize + BoardCellSize/2
	y = BoardY + float32(p.Y)*BoardCellSize + BoardCellSize/2
	return x, y
}

func (g *Game) drawKnightPosition(screen *ebiten.Image, p knight.Position) {
	var xp, yp, strokeWidth, radius float32
	strokeWidth = 1
	radius = 5

	// get current float coordinates of knight
	xp, yp = getCoordinatesFromPosition(p)
	// draw circle
	vector.StrokeCircle(screen,
		xp, yp, radius,
		strokeWidth, color.RGBA{0xaa, 0xaa, 0xaa, 0xff}, false)
}

func (g *Game) drawKnightMove(screen *ebiten.Image, l, p knight.Position) {
	var xp, yp, strokeWidth, xl, yl float32
	strokeWidth = 1

	// get last float coordinates of knight
	xl, yl = getCoordinatesFromPosition(l)
	// get current float coordinates of knight
	xp, yp = getCoordinatesFromPosition(p)
	// draw line between position
	vector.StrokeLine(screen,
		xp, yp, xl, yl,
		strokeWidth, color.RGBA{0xff, 0xff, 0xff, 0xff}, false)
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

func (g *Game) drawFinish(screen *ebiten.Image) {
	if g.state == GameWon && g.blink {
		text.DrawTextAtPos(
			screen, fonts.BigFont,
			55,
			135,
			"SOLUTION\nFOUND",
			color.RGBA{R: 0x0f, G: 0xff, B: 0x0f, A: 0xaf},
		)
	}
	if g.state == GameLost && g.blink {
		text.DrawTextAtPos(
			screen, fonts.BigFont,
			55,
			135,
			"NO\nSOLUTION",
			color.RGBA{R: 0xff, G: 0x0f, B: 0x0f, A: 0xaf},
		)
	}
}
