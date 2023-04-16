package config

import "image/color"

/* Represents the color theme of the game. */
type Theme struct {
	// Background color - essentially, border color
	Background *color.RGBA

	// Color of the alive cell
	CellAlive  *color.RGBA

	// Color of the dead cell
	CellDead   *color.RGBA
}

/* Loads color theme from file. If the file does not exist, returns default theme. */
func LoadTheme() *Theme {
	return LoadDefaultTheme()
}

/* Returns the default color theme. */
func LoadDefaultTheme() *Theme {
	return &Theme{
		Background: &color.RGBA{0x5f, 0x4d, 0x32, 0xff},
		CellAlive : &color.RGBA{0x00, 0x00, 0x00, 0xff},
		CellDead  : &color.RGBA{0xff, 0xf8, 0xf2, 0xff},
	}
}
