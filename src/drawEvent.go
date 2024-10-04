package main

import (
	"github.com/Zedran/life/src/world"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
DrawEvent handles the LMB press on the Map. It allows the user to change the state
of the cell. If the cursor moves while LMB is pressed, all cells in its path are
set to the state opposite to this of a cell originally pressed. This allows the user
to "draw" and "erase" life.
*/
type DrawEvent struct {
	// True indicates that the user is currently drawing world.State on the Map
	Active bool

	// The world.State value that is drawn. Depends on the state of the first cell that was clicked on.
	// If the first cell is alive, all cells in cursor's path are set to world.DEAD.
	// If the first cell is dead, all cells in cursor's path are set to world.ALIVE.
	Draws world.State

	// Current cursor coordinates
	CurX, CurY int
}

/* Returns current coordinates. */
func (de *DrawEvent) Position() (int, int) {
	return de.CurX, de.CurY
}

/* Updates DrawEvent coordinates. Deactivates the event on mouse button up.*/
func (de *DrawEvent) Update() {
	if !de.Active {
		return
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		de.Active = false
		return
	}

	de.CurX, de.CurY = ebiten.CursorPosition()
}

/* Returns new DrawEvent. */
func NewDrawEvent(m *Map) *DrawEvent {
	x, y := ebiten.CursorPosition()

	mx, my := m.GetCellAtPoint(x, y)

	var s world.State

	switch m.World.Cells[my*m.World.Size+mx] {
	case world.ALIVE:
		s = world.DEAD
	case world.DEAD:
		s = world.ALIVE
	}

	return &DrawEvent{
		Active: true,
		Draws:  s,
		CurX:   x,
		CurY:   y,
	}
}
