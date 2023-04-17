package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/world"
)

/* Represents the general structure of the game. */
type Game struct {
	// Game config
	Config *config.Config

	Map    *Map

	// Game logic
	World  *world.World
}

/* Draws interface elements onto the screen. */
func (g *Game) Draw(screen *ebiten.Image) {
	g.Map.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.Config.Window.W), int(g.Config.Window.H)
}

func (g *Game) Update() error {
	g.World.Update()
	return nil
}

/* Creates new game. */
func NewGame() *Game {
	config := config.LoadConfig()

	ebiten.SetWindowSize(int(config.Window.W), int(config.Window.H))
	ebiten.SetWindowTitle(config.Language.Title)

	world := world.Genesis(uint64(config.WorldSize))	

	g := Game{
		Config: config,
		Map   : NewMap(config.Window.W, config.Window.H, config.Theme, world),
		World : world,
	}

	return &g
}
