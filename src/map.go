package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/config/theme"
	"github.com/Zedran/life/src/world"
)

// Size of the border between cells [px]
const BORDER_SIZE float32 =  1

/* Represents graphical world of the game. */
type Map struct {
	// Cached image of the map with all the cells dead. Alive cells are drawn on top. 
	// Creating the whole map from scratch every frame significantly hinders performance.
	Background *ebiten.Image

	// Image of alive cell at maximum zoom
	AliveImg   *ebiten.Image

	// Image of dead cell at maximum zoom
	DeadImg    *ebiten.Image

	// A fraction of maximum cell size by which cell images must be transformed 
	// to fit the grid at current zoom level
	CellScale  float64

	// Can either be equal to 0 or BORDER_SIZE
	BorderSize float32

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
	Theme      *theme.MapTheme

	// A pointer to the logical world of the game
	World      *world.World

	// Current zoom of the map, expressed as index on Map.ZoomSteps slice.
	Zoom       int
	
	// Valid zoom steps for a map of specified size
	// CellSize = Map.ZoomSteps[MapZoom] - BORDER_SIZE
	ZoomSteps  []float32
}

/*
   Adjusts zoom level, moving up or down the Map.ZoomSteps slice. zoom = cellSize + BORDER_SIZE.
   Direction equal to 0 indicates that this function was called by Map constructor.
*/
func (m *Map) AdjustZoomLevel(direction int) {
	if direction == 0 {
		m.Zoom = Index(m.ZoomSteps, GetClosestToMean(m.ZoomSteps))

		m.RowLength  = m.WindowW / m.GetCurrentZoom()
		m.ColHeight  = m.WindowH / m.GetCurrentZoom()
		
		m.RecalculateCellScale()
		
		m.CreateBackground()
		return
	}

	if m.Zoom + direction < 0 || m.Zoom + direction >= len(m.ZoomSteps) {
		return
	}

	cursorX, cursorY := ebiten.CursorPosition()
	oldX, oldY       := m.GetCellAtPoint(cursorX, cursorY)

	m.Zoom += direction

	m.RowLength  = m.WindowW / m.GetCurrentZoom()
	m.ColHeight  = m.WindowH / m.GetCurrentZoom()

	cursorX, cursorY = ebiten.CursorPosition()
	newX, newY       := m.GetCellAtPoint(cursorX, cursorY)

	dX := -float32(newX - oldX)
	dY := -float32(newY - oldY)

	m.Move(dX, dY)

	m.RecalculateCellScale()

	m.CreateBackground()
}

/* Moves the view to the center of the map. */
func (m *Map) CenterView() {
	m.OffSetX = (float32(m.World.Size) - m.RowLength) / 2
	m.OffSetY = (float32(m.World.Size) - m.ColHeight) / 2
}

/* Draws the empty cell map into Map.Background. */
func (m *Map) CreateBackground() {
	m.Background.Fill(m.Theme.Background)

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(m.CellScale, m.CellScale)

	for y := float32(0); y < m.ColHeight; y++ {
		for x := float32(0); x < m.RowLength; x++ {
			m.Background.DrawImage(m.DeadImg, op)

			op.GeoM.Translate(float64(m.GetCurrentZoom()), 0)
		}
		op.GeoM.Translate(-float64(m.RowLength * m.GetCurrentZoom()), float64(m.GetCurrentZoom()))
	}
}

/* Used to fill cell images. */
func (m *Map) CreateCellImages() {
	m.AliveImg.Fill(m.Theme.CellAlive)
	m.DeadImg.Fill(m.Theme.CellDead)
}

/* Draws the cached image of the empty map (Map.Background) to the screen and inserts alive cells on top of it. */
func (m *Map) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.Background, nil)

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(m.CellScale, m.CellScale)

	for y := float32(0); y < m.ColHeight; y++ {
		for x := float32(0); x < m.RowLength; x++ {

			if m.World.Cells[int(y + m.OffSetY) * m.World.Size + int(x + m.OffSetX)] == world.ALIVE {
				screen.DrawImage(m.AliveImg, op)
			}

			op.GeoM.Translate(float64(m.GetCurrentZoom()), 0)
		}
		op.GeoM.Translate(-float64(m.RowLength * m.GetCurrentZoom()), float64(m.GetCurrentZoom()))
	}
}

/* Accepts a screen position in pixels and returns coordinates (x, y) of a cell at this position. */
func (m *Map) GetCellAtPoint(pX, pY int) (x, y int) {
	x = int(float32(pX) * m.RowLength / m.WindowW + m.OffSetX)
	y = int(float32(pY) * m.ColHeight / m.WindowH + m.OffSetY)
	return 
}

/* Get current zoom level. */
func (m *Map) GetCurrentZoom() float32 {
	return m.ZoomSteps[m.Zoom]
}

/* Sets the state s of the cell at (x, y) coordinates. */
func (m *Map) SetState(x, y int, s world.State) {
	m.World.Cells[y * m.World.Size + x] = s
}

/* Offsets the map by specified number of cells. Does nothing for values that exceed the world bounds. */
func (m *Map) Move(dX, dY float32) {
	if m.OffSetX + dX < 0 {
		m.OffSetX = 0
	} else if m.OffSetX + dX > float32(m.World.Size) - m.RowLength {
		m.OffSetX = float32(m.World.Size) - m.RowLength
	} else {
		m.OffSetX += dX
	}

	if m.OffSetY + dY < 0 {
		m.OffSetY = 0
	} else if m.OffSetY + dY > float32(m.World.Size) - m.ColHeight {
		m.OffSetY = float32(m.World.Size) - m.ColHeight
	} else {
		m.OffSetY += dY
	}
}

/* Calls the Map.Move method after translating the movement from graphical measurements into world dimensions. */
func (m *Map) Pan(dX, dY int) {
	m.Move(
		float32(dX) / m.GetCurrentZoom(), 
		float32(dY) / m.GetCurrentZoom(),
	)
}

/* Recalculates the fraction of maximum cell size by which Map.AliveImg and Map.DeadImg are scaled. */
func (m *Map) RecalculateCellScale() {
	m.CellScale = float64((m.GetCurrentZoom() - m.BorderSize) / float32(m.AliveImg.Bounds().Dx()))
}

/* Called after Game.DragEvent finishes its job to eliminate Map.OffSetX and Map.OffSetY inaccuracies. */
func (m *Map) TruncOffSets() {
	m.OffSetX = float32(math.Trunc(float64(m.OffSetX)))
	m.OffSetY = float32(math.Trunc(float64(m.OffSetY)))
}

/* Creates new graphical map of the world. */
func NewMap(windowWidth, windowHeight float32, theme *theme.MapTheme, world *world.World) *Map {
	var m Map

	if theme.Border == true {
		m.BorderSize = BORDER_SIZE
	} else {
		m.BorderSize = 0
	}

	m.WindowW    = windowWidth
	m.WindowH    = windowHeight

	m.OffSetX    = 0
	m.OffSetY    = 0

	m.Theme      = theme
	m.World      = world

	m.ZoomSteps  = GetCommonDivisors(config.ZOOM_MIN, config.ZOOM_MAX, windowWidth, windowHeight)

	maxTileSize := int(m.ZoomSteps[len(m.ZoomSteps) - 1] - m.BorderSize)

	m.Background = ebiten.NewImage(int(windowWidth), int(windowHeight))
	
	m.AliveImg   = ebiten.NewImage(maxTileSize, maxTileSize)
	m.DeadImg    = ebiten.NewImage(maxTileSize, maxTileSize)

	m.CreateCellImages()

	m.AdjustZoomLevel(0)
	m.CenterView()

	return &m
}
