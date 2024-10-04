package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/* Drag event handles panning of the Game.Map */
type DragEvent struct {
	// Indicates that the map is dragged
	Active bool

	// Current cursor coordinates
	CurX, CurY int

	// Cursor coordinates before dragging
	InitX, InitY int
}

/* Returns the difference between current and initial coordinates. */
func (de *DragEvent) PositionDiff() (int, int) {
	return de.CurX - de.InitX, de.CurY - de.InitY
}

/* Returns current coordinates. */
func (de *DragEvent) Position() (int, int) {
	return de.CurX, de.CurY
}

/* Sets new initial position. It is called after the map was panned by the difference. */
func (de *DragEvent) SetNewInit() {
	de.InitX, de.InitY = ebiten.CursorPosition()
}

/* Updates DragEvent coordinates. Deactivates the event on mouse button up.*/
func (de *DragEvent) Update() {
	if !de.Active {
		return
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		de.Active = false
		return
	}

	de.CurX, de.CurY = ebiten.CursorPosition()
}

/* Returns new DragEvent. */
func NewDragEvent() *DragEvent {
	x, y := ebiten.CursorPosition()

	return &DragEvent{
		Active: true,
		CurX:   x,
		CurY:   y,
		InitX:  x,
		InitY:  y,
	}
}
