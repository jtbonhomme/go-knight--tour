package fonts

import (
	// import embed to load truetype font
	_ "embed"
	"log"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
)

const (
	dpi           float64 = 72
	FontSize      int     = 30
	SmallFontSize int     = 15
)

//go:embed Roboto.ttf
var robotoFontData []byte
var SmallFont font.Face
var DefaultFont font.Face

//var DefaultTitleFont font.Face
//var DefaultDebugFont font.Face

func init() {
	var err error

	robotoFont, err := truetype.Parse(robotoFontData)
	if err != nil {
		log.Fatal(err)
	}

	DefaultFont = truetype.NewFace(robotoFont, &truetype.Options{
		Size:    float64(FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	SmallFont = truetype.NewFace(robotoFont, &truetype.Options{
		Size:    float64(SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
