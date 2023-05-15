package ui

import (
	"github.com/Zedran/life/src/config/lang"
	"github.com/Zedran/life/src/config/theme"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

/*
	Creates the elements of the info panel and adds them to it. Returns values of LabeledDisplays:
		gv - generation value
		sv - speed value
		zv - zoom value
*/
func createInfoElements(theme *theme.UITheme, lang *lang.Language, font *font.Face, info *widget.Container) (gv, sv, zv *widget.Label) {
	genDisplay,  gv := NewLabeledDisplay(theme.Generation, font, lang.Generation)
	spdDisplay,  sv := NewLabeledDisplay(theme.Speed,      font, lang.Speed     )
	zoomDisplay, zv := NewLabeledDisplay(theme.Zoom,       font, lang.Zoom      )

	info.AddChild(genDisplay)
	info.AddChild(spdDisplay)
	info.AddChild(zoomDisplay)

	return gv, sv, zv
}

/* Creates the elements of the info panel and adds them to it. */
func createPanelElements(theme *theme.UITheme, lang *lang.Language, font *font.Face, c *Controller, rules string, panel *widget.Container) {
	gameControlCluster := NewButtonCluster()
	gameControlCluster.AddChild(NewButton(theme.PlayToggle,  font, ICON_PLAY_TOGGLE,  c, PLAY_TOGGLE ))
	gameControlCluster.AddChild(NewButton(theme.SlowDown,    font, ICON_SLOW_DOWN,    c, SLOW_DOWN   ))
	gameControlCluster.AddChild(NewButton(theme.SpeedUp,     font, ICON_SPEED_UP,     c, SPEED_UP    ))

	fillControlCluster := NewButtonCluster()
	fillControlCluster.AddChild(NewButton(theme.ResetState,  font, ICON_RESET_STATE,  c, RESET_STATE ))
	fillControlCluster.AddChild(NewButton(theme.RandomState, font, ICON_RANDOM_STATE, c, RANDOM_STATE))

	jumpControlCluster := NewButtonCluster()
	jumpControlCluster.AddChild(NewButton(theme.FF_I,        font, ICON_FF_I,         c, FF_I        ))
	jumpControlCluster.AddChild(NewButton(theme.FF_X,        font, ICON_FF_X,         c, FF_X        ))
	jumpControlCluster.AddChild(NewButton(theme.FF_L,        font, ICON_FF_L,         c, FF_L        ))
	jumpControlCluster.AddChild(NewButton(theme.FF_C,        font, ICON_FF_C,         c, FF_C        ))
	jumpControlCluster.AddChild(NewButton(theme.FF_M,        font, ICON_FF_M,         c, FF_M        ))

	panel.AddChild(gameControlCluster)
	panel.AddChild(fillControlCluster)
	panel.AddChild(jumpControlCluster)

	panel.AddChild(NewTextInput(theme.Rules, font, lang.Rules, rules, c, NEW_RULES))
}
