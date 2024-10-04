package ui

import (
	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

func loadFont(size float64) (font.Face, error) {
	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttf,
		&truetype.Options{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		},
	), nil
}

func loadMonoFont(size float64) (font.Face, error) {
	ttf, err := truetype.Parse(gomono.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttf,
		&truetype.Options{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		},
	), nil
}
