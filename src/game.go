package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/ui"
	"github.com/Zedran/life/src/world"
)

/* Represents the general structure of the game. */
type Game struct {
	// Game config
	Config    *config.Config

	// Drag event handles panning of the Game.Map
	DragEvent *DragEvent

	// Drag event handles drawing states on the Game.Map (LMB + move)
	DrawEvent *DrawEvent

	// Generations clock manages the speed of transition through generations
	GenClock  *Clock

	Map       *Map

	// A pointer to user interface of the game
	UI        *ui.UI

	// Game logic
	World     *world.World

	// Indicates whether the game is currently running
	State     GameState
}

/* Returns true if cursor is within bounds of the main window. */
func (g *Game) CursorInBounds() bool {
	x, y := ebiten.CursorPosition()

	return (x >= 0 && x < int(g.Config.Window.W)) && (y >= 0 && y < int(g.Config.Window.H))
}

/* Draws the map and interface elements onto the screen. */
func (g *Game) Draw(screen *ebiten.Image) {
	g.Map.Draw(screen)
	g.UI.Draw(screen)
}

/* Handles the input from g.UI. */
func (g *Game) HandleControllerInput(uiResp *ui.UIResponse) {
	if uiResp == nil {
		return
	}

	switch uiResp.Signal {
	case ui.PLAY_TOGGLE:
		if g.State == RUN {
			g.State = PAUSE
		} else if g.State == PAUSE {
			g.State = RUN
		}
	case ui.SLOW_DOWN:
		g.GenClock.AdjustSpeed(-1)
		g.UI.UpdateSpeedValue(g.GenClock.GetEventsPerSec())
	case ui.SPEED_UP:
		g.GenClock.AdjustSpeed(1)
		g.UI.UpdateSpeedValue(g.GenClock.GetEventsPerSec())
	case ui.RESET_STATE:
		g.World.Reset()
		g.State = PAUSE
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.RANDOM_STATE:
		g.World.RandomState(5)
		g.State = PAUSE
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.FF_I:
		g.World.Update()
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.FF_X:
		g.World.UpdateBy(10)
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.FF_L:
		g.World.UpdateBy(50)
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.FF_C:
		g.World.UpdateBy(100)
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.FF_M:
		g.World.UpdateBy(1000)
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.NEW_RULES:
		g.World.Rules, _ = world.NewRules(uiResp.Rules)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.Config.Window.W), int(g.Config.Window.H)
}

/* Updates the game every tick. */
func (g *Game) Update() error {
	if !g.UI.Clicked() {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			g.DragEvent = NewDragEvent()
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.DrawEvent = NewDrawEvent(g.Map)	
		}
	}

	if g.DragEvent != nil {
		g.UpdateDragEvent()
	}

	if g.DrawEvent != nil {
		g.UpdateDrawEvent()
	}

	_, dy := ebiten.Wheel()

	if dy > 0 {
		g.Map.AdjustZoomLevel(1)
		g.UI.UpdateZoomValue(g.Map.GetCurrentZoom())
	} else if dy < 0 {
		g.Map.AdjustZoomLevel(-1)
		g.UI.UpdateZoomValue(g.Map.GetCurrentZoom())
	}

	g.HandleControllerInput(g.UI.Update())

	if g.State == RUN && g.GenClock.Tick() == TRIGGER {
		g.World.Update()
		g.UI.UpdateGenValue(g.World.Generation)
	}

	return nil
}

/* Updates drag event and pans the Game.Map accordingly. */
func (g *Game) UpdateDragEvent() {
	g.DragEvent.Update()

	if !g.DragEvent.Active {
		g.Map.TruncOffSets()
		g.DragEvent = nil
		return
	}

	dX, dY := g.DragEvent.PositionDiff()

	g.Map.Pan(-dX, -dY)

	g.DragEvent.SetNewInit()
}

/* Updates g.DrawEvent and sets the state of cell at which the cursor is currently pointing. */
func (g *Game) UpdateDrawEvent() {
	g.DrawEvent.Update()

	if !g.DrawEvent.Active {
		g.DrawEvent = nil
		return
	}

	if g.CursorInBounds() && !g.UI.Clicked() {
		x, y := g.Map.GetCellAtPoint(g.DrawEvent.Position())
		g.Map.SetState(x, y, g.DrawEvent.Draws)
	}
}

/* Creates new game. */
func NewGame() *Game {
	config := config.LoadConfig()

	ebiten.SetWindowSize(int(config.Window.W), int(config.Window.H))
	ebiten.SetWindowTitle(config.Language.Title)

	w := world.Genesis(config.WorldSize)
	
	ui, err := ui.NewUI(world.DEFAULT_RULES)
	if err != nil {
		log.Fatal(err)
	}

	g := Game{
		Config   : config,
		DragEvent: nil,
		DrawEvent: nil,
		GenClock : NewClock(),
		Map      : NewMap(config.Window.W, config.Window.H, config.Theme, w),
		UI       : ui,
		World    : w,
		State    : PAUSE,
	}

	g.UI.UpdateSpeedValue(g.GenClock.GetEventsPerSec())
	g.UI.UpdateZoomValue(g.Map.GetCurrentZoom())
	g.UI.UpdateGenValue(0)

	return &g
}
