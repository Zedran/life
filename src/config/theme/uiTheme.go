package theme

import "image/color"

/* User Interface theme. */
type UITheme struct {
	InfoPanel *PanelTheme
	CtrlPanel *PanelTheme

	Generation *LabelTheme
	Speed      *LabelTheme
	Zoom       *LabelTheme

	PlayToggle *ButtonTheme
	SlowDown   *ButtonTheme
	SpeedUp    *ButtonTheme

	ResetState  *ButtonTheme
	RandomState *ButtonTheme

	FF_I *ButtonTheme
	FF_X *ButtonTheme
	FF_L *ButtonTheme
	FF_C *ButtonTheme
	FF_M *ButtonTheme

	Rules *TextInputTheme
}

type ButtonTheme struct {
	Text *color.NRGBA

	Idle    *color.NRGBA
	Hover   *color.NRGBA
	Pressed *color.NRGBA
}

type LabelTheme struct {
	Text *color.RGBA
}

type PanelTheme struct {
	Background *color.NRGBA
}

type TextInputTheme struct {
	Text          *color.RGBA
	Idle          *color.RGBA
	Disabled      *color.RGBA
	Caret         *color.RGBA
	DisabledCaret *color.RGBA
}
