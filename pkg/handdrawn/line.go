package handdrawn

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Roughness float32 = 1.2
)

var (
	r             *rand.Rand
	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomize(f float32) float32 {
	// Float32 and Float64 values are in [0, 1).
	return f - Roughness/2 + Roughness + r.Float32()

}

func Line(screen *ebiten.Image,
	x1, y1, x2, y2, strokeWidth float32,
	clr color.Color,
	antialias bool) {
	randomizedLine(screen,
		x1, y1, x2, y2,
		strokeWidth, clr, antialias)
}

func randomizedLine(screen *ebiten.Image,
	x1, y1, x2, y2, strokeWidth float32,
	clr color.Color,
	antialias bool) {
	// randomness for start and end points
	x1 = randomize(x1)
	y1 = randomize(y1)
	x2 = randomize(x2)
	y2 = randomize(y2)

	// bezier
	var path vector.Path

	// curve origin
	path.MoveTo(x1, y1)

	// add next point with vectors
	cpx0, cpy0 := x1, y1
	cpx1, cpy1 := x2, y2
	cpx0 += 30
	cpx1 -= 30
	path.CubicTo(cpx0, cpy0, cpx1, cpy1, x2, y2)

	var vs []ebiten.Vertex
	var is []uint16
	opl := &vector.StrokeOptions{}
	opl.Width = strokeWidth
	opl.LineJoin = vector.LineJoinRound
	vs, is = path.AppendVerticesAndIndicesForStroke(nil, nil, opl)

	R, G, B, _ := clr.RGBA()
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(R) / float32(0xff)
		vs[i].ColorG = float32(G) / float32(0xff)
		vs[i].ColorB = float32(B) / float32(0xff)
		vs[i].ColorA = 1
	}

	opt := &ebiten.DrawTrianglesOptions{}
	opt.AntiAlias = antialias
	opt.FillRule = ebiten.NonZero
	screen.DrawTriangles(vs, is, whiteSubImage, opt)
}

func randomizedStrokeLine(screen *ebiten.Image,
	x1, y1, x2, y2, strokeWidth float32,
	clr color.Color,
	antialias bool) {
	// randomness for start and end points
	x1 = randomize(x1)
	y1 = randomize(y1)
	x2 = randomize(x2)
	y2 = randomize(y2)

	vector.StrokeLine(screen,
		x1, y1, x2, y2,
		strokeWidth, clr, antialias)
}
