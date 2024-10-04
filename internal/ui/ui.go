package ui

import (
	"fmt"

	"github.com/Zedran/life/internal/config/lang"
	"github.com/Zedran/life/internal/config/theme"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

/* Top structure of User Interface. */
type UI struct {
	// Controller shared by widgets
	controller *Controller

	// UI struct provided by ebitenui
	ui *ebitenui.UI

	// Root container of the UI
	root *widget.Container

	// Info panel
	info *widget.Container

	// Control panel
	panel *widget.Container

	// Generation
	genValue *widget.Label

	// Current zoom
	zoomValue *widget.Label

	// Current speed
	speedValue *widget.Label

	// TextInput for setting the rules
	rules *widget.TextInput
}

/* Draws UI onto the screen. */
func (ui *UI) Draw(screen *ebiten.Image) {
	ui.ui.Draw(screen)
}

/* Called on click. Returns true if a mouse cursor is inside one of the panels. */
func (ui *UI) Clicked() bool {
	x, y := ebiten.CursorPosition()

	for _, c := range []*widget.Container{ui.info, ui.panel} {
		r := c.GetWidget().Rect

		if (x >= r.Min.X && x <= r.Max.X) && (y >= r.Min.Y && y <= r.Max.Y) {
			return true
		}
	}

	return false
}

/* Sets the contents of the ui.rules. If user specifies incorrect rules, this method is called to revert to the previous ones. */
func (ui *UI) SetRules(rules string) {
	ui.rules.InputText = rules
}

/* Updates the UI. Returns UIResponse if widget was interacted with. */
func (ui *UI) Update() *UIResponse {
	ui.ui.Update()

	if ui.controller.GetSignal() != NONE {
		return ui.controller.Emit()
	}

	return nil
}

/* Sets the generation label text to new value. */
func (ui *UI) UpdateGenValue(new uint64) {
	ui.genValue.Label = fmt.Sprintf("%-6d", new)
}

/* Sets the current speed label text to new value. */
func (ui *UI) UpdateSpeedValue(new int) {
	ui.speedValue.Label = fmt.Sprintf("%-3d", new)
}

/* Sets the current zoom label text to new value. */
func (ui *UI) UpdateZoomValue(new float32) {
	ui.zoomValue.Label = fmt.Sprintf("%-3.0f", new)
}

/* Creates UI elements: containers, widgets and the corresponding Controller. */
func NewUI(uit *theme.UITheme, lang *lang.Language, rules string) (*UI, error) {
	root := createRoot()
	spacer := createSpacer()
	info := createInfo(uit)
	panel := createPanel(uit)

	font, err := loadFont(15)
	if err != nil {
		return nil, err
	}

	monoFont, err := loadMonoFont(15)
	if err != nil {
		return nil, err
	}

	c := NewController(rules)

	gv, sv, zv := createInfoElements(uit, lang, &monoFont, info)

	rulesTI := createPanelElements(uit, lang, &font, c, rules, panel)

	root.AddChild(spacer)
	root.AddChild(info)
	root.AddChild(panel)

	eui := ebitenui.UI{
		Container: root,
	}

	return &UI{
		controller: c,
		ui:         &eui,
		root:       root,
		info:       info,
		panel:      panel,
		genValue:   gv,
		zoomValue:  zv,
		speedValue: sv,
		rules:      rulesTI,
	}, nil
}
