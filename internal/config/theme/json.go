package theme

import (
	"image/color"
)

/* A JSON representation of Theme. */
type jsonTheme struct {
	Map jsonMapTheme `json:"map"`

	InfoPanel jsonPanelTheme `json:"info_panel"`
	CtrlPanel jsonPanelTheme `json:"ctrl_panel"`

	Generation jsonLabelTheme `json:"generation"`
	Speed      jsonLabelTheme `json:"speed"`
	Zoom       jsonLabelTheme `json:"zoom"`

	PlayToggle jsonButtonTheme `json:"play_toggle"`
	SlowDown   jsonButtonTheme `json:"slow_down"`
	SpeedUp    jsonButtonTheme `json:"speed_up"`

	ResetState  jsonButtonTheme `json:"reset_state"`
	RandomState jsonButtonTheme `json:"random_state"`

	FF_I jsonButtonTheme `json:"ff_1"`
	FF_X jsonButtonTheme `json:"ff_10"`
	FF_L jsonButtonTheme `json:"ff_50"`
	FF_C jsonButtonTheme `json:"ff_100"`
	FF_M jsonButtonTheme `json:"ff_1000"`

	Rules jsonTextInputTheme `json:"rules"`
}

/*
Converts JSON struct into Theme. If data structure of any of the elements is invalid,
it is replaced with a default configuration.
*/
func (t *jsonTheme) ToTheme() *Theme {
	return &Theme{
		MapTheme: t.Map.ToTheme(),
		UITheme: &UITheme{
			InfoPanel:   t.InfoPanel.ToTheme(),
			CtrlPanel:   t.CtrlPanel.ToTheme(),
			Generation:  t.Generation.ToTheme(),
			Speed:       t.Speed.ToTheme(),
			Zoom:        t.Zoom.ToTheme(),
			PlayToggle:  t.PlayToggle.ToTheme(),
			SlowDown:    t.SlowDown.ToTheme(),
			SpeedUp:     t.SpeedUp.ToTheme(),
			ResetState:  t.ResetState.ToTheme(),
			RandomState: t.RandomState.ToTheme(),
			FF_I:        t.FF_I.ToTheme(),
			FF_X:        t.FF_X.ToTheme(),
			FF_L:        t.FF_L.ToTheme(),
			FF_C:        t.FF_C.ToTheme(),
			FF_M:        t.FF_M.ToTheme(),
			Rules:       t.Rules.ToTheme(),
		},
	}
}

/* A JSON representation of ButtonTheme. */
type jsonButtonTheme struct {
	Text    string `json:"text"`
	Idle    string `json:"idle"`
	Hover   string `json:"hover"`
	Pressed string `json:"pressed"`
}

/* Converts JSON struct into ButtonTheme. If data structure is invalid, returns default configuration. */
func (jbt *jsonButtonTheme) ToTheme() *ButtonTheme {
	vals := make([]*color.NRGBA, 4)

	for i, v := range []string{jbt.Text, jbt.Idle, jbt.Hover, jbt.Pressed} {
		c, err := decodeColor(v)
		if err != nil {
			return defaultButtonTheme.ToTheme()
		}

		vals[i] = rgbaToNRGBA(c)
		if err != nil {
			return defaultButtonTheme.ToTheme()
		}
	}

	return &ButtonTheme{
		Text:    vals[0],
		Idle:    vals[1],
		Hover:   vals[2],
		Pressed: vals[3],
	}
}

/* A JSON representation of LabelTheme. */
type jsonLabelTheme struct {
	Text string `json:"text"`
}

/* Converts JSON struct into LabelTheme. If data structure is invalid, returns default configuration. */
func (jlt *jsonLabelTheme) ToTheme() *LabelTheme {
	text, err := decodeColor(jlt.Text)
	if err != nil {
		return defaultLabelTheme.ToTheme()
	}

	return &LabelTheme{
		Text: text,
	}
}

/* A JSON representation of MapTheme. */
type jsonMapTheme struct {
	Border     bool   `json:"border"`
	Background string `json:"background"`
	CellAlive  string `json:"cell_alive"`
	CellDead   string `json:"cell_dead"`
}

/* Converts JSON struct into MapTheme. If data structure is invalid, returns default configuration. */
func (jmt *jsonMapTheme) ToTheme() *MapTheme {
	var (
		vals = make([]*color.RGBA, 3)
		err  error
	)

	for i, v := range []string{jmt.Background, jmt.CellAlive, jmt.CellDead} {
		vals[i], err = decodeColor(v)
		if err != nil {
			return defaultMapTheme.ToTheme()
		}
	}

	return &MapTheme{
		Border:     jmt.Border,
		Background: vals[0],
		CellAlive:  vals[1],
		CellDead:   vals[2],
	}
}

/* A JSON representation of PanelTheme. */
type jsonPanelTheme struct {
	Background string `json:"background"`
}

/* Converts JSON struct into PanelTheme. If data structure is invalid, returns default configuration. */
func (jpt *jsonPanelTheme) ToTheme() *PanelTheme {
	bg, err := decodeColor(jpt.Background)
	if err != nil {
		return defaultPanelTheme.ToTheme()
	}

	return &PanelTheme{
		Background: rgbaToNRGBA(bg),
	}
}

/* A JSON representation of TextInputTheme. */
type jsonTextInputTheme struct {
	Text          string `json:"text"`
	Idle          string `json:"idle"`
	Disabled      string `json:"disabled"`
	Caret         string `json:"caret"`
	DisabledCaret string `json:"disabled_caret"`
}

/* Converts JSON struct into TextInputTheme. If data structure is invalid, returns default configuration. */
func (jtit *jsonTextInputTheme) ToTheme() *TextInputTheme {
	var (
		vals = make([]*color.RGBA, 5)
		err  error
	)

	for i, v := range []string{jtit.Text, jtit.Idle, jtit.Disabled, jtit.Caret, jtit.DisabledCaret} {
		vals[i], err = decodeColor(v)
		if err != nil {
			return defaultTextInputTheme.ToTheme()
		}
	}

	return &TextInputTheme{
		Text:          vals[0],
		Idle:          vals[1],
		Disabled:      vals[2],
		Caret:         vals[3],
		DisabledCaret: vals[4],
	}
}
