package theme

import "image/color"

type MapTheme struct {
	// If true, a border will be present around cells
	Border     bool

	// Background color - essentially, border color
	Background *color.RGBA

	// Color of the alive cell
	CellAlive  *color.RGBA

	// Color of the dead cell
	CellDead   *color.RGBA
}
