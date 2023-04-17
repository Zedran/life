package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/world"
)

const (
	// Initial zoom level
	ZOOM_INIT   float32 =  8

	// Size of the border between cells in pixels
	BORDER_SIZE float32 =  1
)

/* Represents graphical world of the game. */
type Map struct {
	// Cached image of the map with all the cells dead. Alive cells are drawn on top. 
	// Creating the whole map from scratch every frame significantly hinders performance.
	Background *ebiten.Image

	// Number of cells displayed in one row
	RowLength  float32

	// Number of cells displayed in one column
	ColHeight  float32

	// Offsets are not pixel values, they are counted in relation to world.Cells
	OffSetX    float32
	OffSetY    float32

	// Window width
	WindowW    float32

	// Window height
	WindowH    float32

	// A pointer to color theme of the game
	Theme      *config.Theme

	// A pointer to the logical world of the game
	World      *world.World

	// Current zoom of the map. CellSize = Game.Zoom - BORDER_SIZE
	Zoom       float32
}

/* Draws the empty cell map into Map.Background. */
func (m *Map) CreateBackground() {
	m.Background.Fill(m.Theme.Background)

	cellSize := m.Zoom - BORDER_SIZE

	for y := float32(0); y < m.ColHeight; y++ {
		for x := float32(0); x < m.RowLength; x++ {
			var color *color.RGBA

			switch m.World.Cells[int(y + m.OffSetY) * m.World.Size + int(x + m.OffSetX)] {
			case world.ALIVE:
				color = m.Theme.CellAlive
			case world.DEAD:
				color = m.Theme.CellDead
			}

			vector.DrawFilledRect(
				m.Background, 
				x * m.Zoom, 
				y * m.Zoom, 
				cellSize, 
				cellSize, 
				color, 
				false,
			)
		}
	}
}

/* Draws the cached image of the empty map (Map.Background) to the screen and inserts alive cells on top of it. */
func (m *Map) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.Background, nil)

	cellSize := m.Zoom - BORDER_SIZE

	for y := float32(0); y < m.ColHeight; y++ {
		for x := float32(0); x < m.RowLength; x++ {

			if m.World.Cells[int(y + m.OffSetY) * m.World.Size + int(x + m.OffSetX)] == world.ALIVE {
				vector.DrawFilledRect(
					screen, 
					x * m.Zoom, 
					y * m.Zoom, 
					cellSize, 
					cellSize, 
					m.Theme.CellAlive, 
					false,
				)
			}
		}
	}
}

/* Sets zoom level in pixels. zoom = cellSize + BORDER_SIZE */
func (m *Map) SetZoomLevel(zoom float32) {
	m.Zoom = zoom

	m.RowLength  = m.WindowW / m.Zoom
	m.ColHeight  = m.WindowH / m.Zoom

	m.CreateBackground()
}

/* Creates new graphical map of the world. */
func NewMap(windowWidth, windowHeight float32, theme *config.Theme, world *world.World) *Map {
	var m Map

	m.WindowW    = windowWidth
	m.WindowH    = windowHeight

	m.OffSetX    = 0
	m.OffSetY    = 0

	m.Theme      = theme
	m.World      = world

	m.Background = ebiten.NewImage(int(windowWidth), int(windowHeight))

	m.SetZoomLevel(ZOOM_INIT)

	return &m
}
