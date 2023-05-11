package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

/* Top structure of User Interface. */
type UI struct {
	// Controller shared by widgets
	controller *Controller

	// UI struct provided by ebitenui
	ui         *ebitenui.UI

	// Root container of the UI
	rootCnt    *widget.Container

	// Lower panel of the UI
	lowerCnt   *widget.Container
}

/* Draws UI onto the screen. */
func (ui *UI) Draw(screen *ebiten.Image) {
	ui.ui.Draw(screen)
}

func (ui *UI) Clicked() bool {
	x, y := ebiten.CursorPosition()
	
	pr := ui.lowerCnt.GetWidget().Rect

	return (x >= pr.Min.X && x <= pr.Max.X) && (y >= pr.Min.Y && y <= pr.Max.Y)
}

/* Updates the UI. Returns UIResponse if widget was interacted with. */
func (ui *UI) Update() *UIResponse {
	ui.ui.Update()

	if ui.controller.GetSignal() != NONE {
		return ui.controller.Emit()
	}

	return nil
}

/* Creates UI elements: containers, widgets and the corresponding Controller. */
func NewUI(rules string) (*UI, error) {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
			),
		),
	)

	panel := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Padding(
					widget.Insets{
						Left:   10,
						Right:  10,
						Top:    10,
						Bottom: 10,
					},
				),
				widget.GridLayoutOpts.Columns(12),
				widget.GridLayoutOpts.Spacing(5, 5),
			),
		),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{150, 150, 150, 150})),

		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionEnd,
					StretchHorizontal:  true,
				},
			),
		),
	)

	font, err := loadFont(20)
	if err != nil {
		return nil, err
	}

	c := NewController(rules)

	panel.AddChild(NewButton(&font, string([]rune{0x25ba, 0x007c}), c, PLAY_TOGGLE))
	panel.AddChild(NewButton(&font, "-", c, SLOW_DOWN))
	panel.AddChild(NewButton(&font, "+", c, SPEED_UP))
	panel.AddChild(NewButton(&font, string(0x00d8), c, RESET_STATE))
	panel.AddChild(NewButton(&font, string([]rune{0x2591, 0x2591}), c, RANDOM_STATE))
	panel.AddChild(NewButton(&font, "I", c, FF_I))
	panel.AddChild(NewButton(&font, "X", c, FF_X))
	panel.AddChild(NewButton(&font, "L", c, FF_L))
	panel.AddChild(NewButton(&font, "C", c, FF_C))
	panel.AddChild(NewButton(&font, "M", c, FF_M))
	panel.AddChild(NewTextInput(&font, "Rules", rules, c, NEW_RULES))

	root.AddChild(panel)

	eui := ebitenui.UI{
		Container: root,
	}

	return &UI{
		controller: c,
		ui        : &eui,
		rootCnt   : root,
		lowerCnt  : panel,
	}, nil
}
