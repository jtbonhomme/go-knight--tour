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
	Roughness float32 = 1.07
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
	randomizedCubicLine(screen,
		x1, y1, x2, y2,
		strokeWidth, clr, antialias)
	randomizedCubicLine(screen,
		x2, y2, x1, y1,
		strokeWidth, clr, antialias)
}

func signMult(n float32) float32 {
	if n <= 0 {
		return -1
	}
	return 1
}

func randomizedCubicLine(screen *ebiten.Image,
	x1, y1, x4, y4, strokeWidth float32,
	clr color.Color,
	antialias bool) {

	// bezier
	var path vector.Path

	// curve origin
	path.MoveTo(randomize(x1), randomize(y1))

	// Line analysis
	slope := (y4 - y1) / (x4 - x1)

	// 1st segment; add next point with vectors
	x2 := x1 + (x4-x1)/2
	y2 := y1 + (y4-y1)/2
	cpx1, cpy1 := x1, y1
	cpx2, cpy2 := x2, y2
	if slope <= 0 {
		cpx1 += signMult((x4 - x1)) * Roughness
		cpy2 += signMult((x4 - x1)) * Roughness

	} else {
		cpx1 += signMult((x4 - x1)) * Roughness
		cpy2 += -signMult((x4 - x1)) * Roughness
	}
	// randomness for start and end points
	path.CubicTo(randomize(cpx1), randomize(cpy1), randomize(cpx2), randomize(cpy2), randomize(x2), randomize(y2))

	// 2nd segment; add next point with vectors
	x3 := x1 + 3*(x4-x1)/4
	y3 := y1 + 3*(y4-y1)/4
	cpx2, cpy2 = x2, y2
	cpx3, cpy3 := x3, y3
	if slope <= 0 {
		cpy2 += -signMult((x4 - x1)) * Roughness
		cpx3 += -signMult((x4 - x1)) * Roughness

	} else {
		cpy2 += signMult((x4 - x1)) * Roughness
		cpx3 += -signMult((x4 - x1)) * Roughness
	}
	// randomness for start and end points
	path.CubicTo(randomize(cpx2), randomize(cpy2), randomize(cpx3), randomize(cpy3), randomize(x3), randomize(y3))

	// 3rd segment; add next point with vectors
	cpx3, cpy3 = x3, y3
	cpx4, cpy4 := x4, y4
	if slope <= 0 {
		cpy3 += signMult((x4 - x1)) * Roughness
		cpx4 += signMult((x4 - x1)) * Roughness

	} else {
		cpy3 += -signMult((x4 - x1)) * Roughness
		cpx4 += signMult((x4 - x1)) * Roughness
	}
	// randomness for start and end points
	path.CubicTo(randomize(cpx3), randomize(cpy3), randomize(cpx4), randomize(cpy4), randomize(x4), randomize(y4))

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
	vector.StrokeLine(screen,
		randomize(x1), randomize(y1), randomize(x2), randomize(y2),
		strokeWidth, clr, antialias)
}
