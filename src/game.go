package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/Zedran/life/src/config"
	"github.com/Zedran/life/src/ui"
	"github.com/Zedran/life/src/world"
)

/* Represents the general structure of the game. */
type Game struct {
	// Game config
	Config *config.Config

	// Drag event handles panning of the Game.Map
	DragEvent *DragEvent

	// Drag event handles drawing states on the Game.Map (LMB + move)
	DrawEvent *DrawEvent

	// Generations clock manages the speed of transition through generations
	GenClock *Clock

	Map *Map

	// A pointer to user interface of the game
	UI *ui.UI

	// Game logic
	World *world.World

	// Sentinel channel used for stopping long tasks
	Sentinel chan bool

	// Indicates whether the game is currently running
	State GameState
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

/* Performs a jump in time by specified number of generations. */
func (g *Game) FastForward(generations int) {
	if g.State == FF {
		return
	}

	prevState := g.State

	g.State = PAUSE

	for g.World.Working {
		// Hold while World.wg finishes its Wait
		time.Sleep(time.Microsecond)
	}

	g.State = FF

	g.World.UpdateBy(&g.Sentinel, generations)

	g.UI.UpdateGenValue(g.World.Generation)

	if prevState == RUN {
		g.State = PAUSE
	} else {
		g.State = prevState
	}
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
		} else {
			g.StopTask()
		}
	case ui.SLOW_DOWN:
		g.GenClock.AdjustSpeed(-1)
		g.UI.UpdateSpeedValue(g.GenClock.GetEventsPerSec())
	case ui.SPEED_UP:
		g.GenClock.AdjustSpeed(1)
		g.UI.UpdateSpeedValue(g.GenClock.GetEventsPerSec())
	case ui.RESET_STATE:
		g.StopTask()
		g.World.Reset()
		g.State = PAUSE
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.RANDOM_STATE:
		g.StopTask()
		g.World.RandomState(5)
		g.State = PAUSE
		g.UI.UpdateGenValue(g.World.Generation)
	case ui.FF_I:
		go g.FastForward(1)
	case ui.FF_X:
		go g.FastForward(10)
	case ui.FF_L:
		go g.FastForward(50)
	case ui.FF_C:
		go g.FastForward(100)
	case ui.FF_M:
		go g.FastForward(1000)
	case ui.NEW_RULES:
		rules, err := world.NewRules(uiResp.Rules)
		if err != nil {
			g.UI.SetRules(g.World.Rules.Str)
			break
		}
		g.World.Rules = rules
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.Config.Window.W), int(g.Config.Window.H)
}

/* Stops currently running long task. Currently, it is only used to stop Game.FastForward. */
func (g *Game) StopTask() {
	if g.State == FF {
		g.Sentinel <- true
	}
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
	} else if g.State == FF {
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
	cfg := config.LoadConfig()

	if err := config.WriteDefaults(); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(int(cfg.Window.W), int(cfg.Window.H))
	ebiten.SetWindowTitle(cfg.Language.Title)
	ebiten.SetScreenClearedEveryFrame(false)

	w := world.Genesis(cfg.WorldSize)

	ui, err := ui.NewUI(cfg.Theme.UITheme, cfg.Language, world.DEFAULT_RULES)
	if err != nil {
		log.Fatal(err)
	}

	g := Game{
		Config:    cfg,
		DragEvent: nil,
		DrawEvent: nil,
		GenClock:  NewClock(),
		Map:       NewMap(cfg, w),
		UI:        ui,
		World:     w,
		State:     PAUSE,
		Sentinel:  make(chan bool),
	}

	g.UI.UpdateSpeedValue(g.GenClock.GetEventsPerSec())
	g.UI.UpdateZoomValue(g.Map.GetCurrentZoom())
	g.UI.UpdateGenValue(0)

	g.World.RandomState(5)

	return &g
}
