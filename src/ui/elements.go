package ui

import (
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/*
	Creates the elements of the info panel and adds them to it. Returns values of LabeledDisplays:
		gv - generation value
		sv - speed value
		zv - zoom value
*/
func createInfoElements(font *font.Face, info *widget.Container) (gv, sv, zv *widget.Label) {
	genDisplay,  gv := NewLabeledDisplay(font, "Generation")
	spdDisplay,  sv := NewLabeledDisplay(font, "Speed")
	zoomDisplay, zv := NewLabeledDisplay(font, "Zoom")

	info.AddChild(genDisplay)
	info.AddChild(spdDisplay)
	info.AddChild(zoomDisplay)

	return gv, sv, zv
}

/* Creates the elements of the info panel and adds them to it. */
func createPanelElements(font *font.Face, c *Controller, rules string, panel *widget.Container) {
	gameControlCluster := NewButtonCluster()
	gameControlCluster.AddChild(NewButton(font, ICON_PLAY_TOGGLE,  c, PLAY_TOGGLE ))
	gameControlCluster.AddChild(NewButton(font, ICON_SLOW_DOWN,    c, SLOW_DOWN   ))
	gameControlCluster.AddChild(NewButton(font, ICON_SPEED_UP,     c, SPEED_UP    ))

	fillControlCluster := NewButtonCluster()
	fillControlCluster.AddChild(NewButton(font, ICON_RESET_STATE,  c, RESET_STATE ))
	fillControlCluster.AddChild(NewButton(font, ICON_RANDOM_STATE, c, RANDOM_STATE))

	jumpControlCluster := NewButtonCluster()
	jumpControlCluster.AddChild(NewButton(font, ICON_FF_I,         c, FF_I        ))
	jumpControlCluster.AddChild(NewButton(font, ICON_FF_X,         c, FF_X        ))
	jumpControlCluster.AddChild(NewButton(font, ICON_FF_L,         c, FF_L        ))
	jumpControlCluster.AddChild(NewButton(font, ICON_FF_C,         c, FF_C        ))
	jumpControlCluster.AddChild(NewButton(font, ICON_FF_M,         c, FF_M        ))

	panel.AddChild(gameControlCluster)
	panel.AddChild(fillControlCluster)
	panel.AddChild(jumpControlCluster)

	panel.AddChild(NewTextInput(font, "Rules", rules, c, NEW_RULES))
}
