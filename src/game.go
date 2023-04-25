package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/world"
)

/* Represents the general structure of the game. */
type Game struct {
	// Game config
	Config    *config.Config

	// Drag event handles panning of the Game.Map
	DragEvent *DragEvent

	// Generations clock manages the speed of transition through generations
	GenClock  *Clock

	Map       *Map

	// Game logic
	World     *world.World
}

/* Draws interface elements onto the screen. */
func (g *Game) Draw(screen *ebiten.Image) {
	g.Map.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.Config.Window.W), int(g.Config.Window.H)
}

/* Updates the game every tick. */
func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.DragEvent = NewDragEvent()
	}

	if g.DragEvent != nil {
		g.UpdateDragEvent()
	}

	_, dy := ebiten.Wheel()

	if dy > 0 {
		g.Map.AdjustZoomLevel(1)
	} else if dy < 0 {
		g.Map.AdjustZoomLevel(-1)
	}

	if g.GenClock.Tick() == TRIGGER {
		g.World.Update()
	}

	return nil
}

/* Updates drag event and pans the Game.Map accordingly. */
func (g *Game) UpdateDragEvent() {
	g.DragEvent.Update()

	if !g.DragEvent.Active {
		g.DragEvent = nil
		return
	}

	dX, dY := g.DragEvent.PositionDiff()

	g.Map.Pan(-dX, -dY)

	g.DragEvent.SetNewInit()
}

/* Creates new game. */
func NewGame() *Game {
	config := config.LoadConfig()

	ebiten.SetWindowSize(int(config.Window.W), int(config.Window.H))
	ebiten.SetWindowTitle(config.Language.Title)

	world := world.Genesis(config.WorldSize)	

	g := Game{
		Config   : config,
		DragEvent: nil,
		GenClock : NewClock(),
		Map      : NewMap(config.Window.W, config.Window.H, config.Theme, world),
		World    : world,
	}

	return &g
}
