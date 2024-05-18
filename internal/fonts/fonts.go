package fonts

import (
	// import embed to load truetype font
	_ "embed"
	"log"

	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font"
)

const (
	dpi             float64 = 72
	DefaultFontSize int     = 30
	SmallFontSize   int     = 15
	BigFontSize     int     = 70
)

//go:embed Roboto.ttf
var robotoFontData []byte
var SmallFont font.Face
var BigFont font.Face
var DefaultFont font.Face

func init() {
	var err error

	robotoFont, err := truetype.Parse(robotoFontData)
	if err != nil {
		log.Fatal(err)
	}

	DefaultFont = truetype.NewFace(robotoFont, &truetype.Options{
		Size:    float64(DefaultFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	SmallFont = truetype.NewFace(robotoFont, &truetype.Options{
		Size:    float64(SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	BigFont = truetype.NewFace(robotoFont, &truetype.Options{
		Size:    float64(BigFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
