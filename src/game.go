package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/world"
)

const (
	ZOOM_MAX    float32 = 20
	ZOOM_INIT   float32 =  8

	BORDER_SIZE float32 =  1
)

/* Represents the general structure of the game. */
type Game struct {
	// Current zoom of the map. CellSize = Game.Zoom - BORDER_SIZE
	Zoom      float32

	// Offsets are not pixel values, they are counted in relation to world.Cells
	OffSetX   float32
	OffSetY   float32

	// Game config
	Config    *config.Config

	// Game logic
	World     *world.World
}

/* Draws interface elements onto the screen. */
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.Config.Theme.Background)

	cellSize := g.Zoom - BORDER_SIZE

	for y := float32(0); y < float32(g.Config.Window.H / g.Zoom); y++ {
		for x := float32(0); x < float32(g.Config.Window.W / g.Zoom); x++ {
			var color *color.RGBA

			switch g.World.Cells[int(y + g.OffSetY) * g.Config.WorldSize + int(x + g.OffSetX)] {
			case world.ALIVE:
				color = g.Config.Theme.CellAlive
			case world.DEAD:
				color = g.Config.Theme.CellDead
			}

			vector.DrawFilledRect(
				screen, 
				x * g.Zoom, 
				y * g.Zoom, 
				cellSize, 
				cellSize, 
				color, 
				false,
			)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.Config.Window.W), int(g.Config.Window.H)
}

func (g *Game) Update() error {
	return nil
}

/* Creates new game. */
func NewGame() *Game {
	config := config.LoadConfig()

	ebiten.SetWindowSize(int(config.Window.W), int(config.Window.H))
	ebiten.SetWindowTitle(config.Language.Title)

	g := Game{
		Zoom   : ZOOM_INIT,
		OffSetX: 0,
		OffSetY: 0,
		Config : config,
		World  : world.Genesis(uint64(config.WorldSize)),
	}

	return &g
}
