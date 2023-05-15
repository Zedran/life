package theme

var (
	defaultButtonTheme = jsonButtonTheme{
		Text   : "#dff4ffff",
		Idle   : "#006e00ff",
		Hover  : "#00aa00ff",
		Pressed: "#005000ff",
	}

	defaultLabelTheme = jsonLabelTheme{
		Text: "#ffffffff",
	}

	defaultMapTheme = jsonMapTheme{
		Background: "#5f4d32ff",
		CellAlive : "#000000ff",
		CellDead  : "#fff8f2ff",
	}

	defaultPanelTheme = jsonPanelTheme{
		Background: "#2b2b2bf0",
	}

	defaultTextInputTheme = jsonTextInputTheme{
		Text         : "#646464ff",
		Idle         : "#ffffffff",
		Disabled     : "#c8c8c8ff",
		Caret        : "#ffffffff",
		DisabledCaret: "#c8c8c8ff",
	}
)
