package theme

import "image/color"

type MapTheme struct {
	// If true, a border will be present around cells. 
	// If config.Config.ZoomMin equals config.ZOOM_MIN_LIMIT, 
	// border is always disabled.
	Border     bool

	// Background color - essentially, border color
	Background *color.RGBA

	// Color of the alive cell
	CellAlive  *color.RGBA

	// Color of the dead cell
	CellDead   *color.RGBA
}
