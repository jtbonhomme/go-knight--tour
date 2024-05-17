package text

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	shadowDist int = 3
)

func DrawTextAtPos(screen *ebiten.Image, f font.Face, x, y int, msg string, c color.Color) {
	// text
	text.Draw(screen,
		msg,
		f,
		x,
		y,
		c)
}

func DrawCenteredText(screen *ebiten.Image, f font.Face, x, y, w, h int, msg string, c color.Color) {
	textDim := text.BoundString(f, msg)

	textWidth := textDim.Max.X - textDim.Min.X
	textHeight := textDim.Max.Y - textDim.Min.Y

	x = x + (w-int(textWidth))/2 + 2
	y = y + (h-int(textHeight))/2 + int(textHeight) + 2

	// text
	text.Draw(screen,
		msg,
		f,
		x,
		y,
		c)
}
