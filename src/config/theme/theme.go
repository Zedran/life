package theme

import (
	"encoding/json"
	"os"
)

/* Represents the color theme of the game. */
type Theme struct {
	// Map theme: cells and border
	MapTheme *MapTheme

	// User interface theme
	UITheme  *UITheme
}

/* Loads color theme from file. If the file does not exist, returns default theme. */
func LoadTheme(path string) *Theme {
	var t jsonTheme

	stream, err := os.ReadFile(path)
	if err != nil {
		return t.ToTheme()
	}


	if err = json.Unmarshal(stream, &t); err != nil {
		return t.ToTheme()
	}

	return t.ToTheme()
}

/* Saves the unexported default theme data. */
func SaveDefault(path string) error {
	pt  := defaultPanelTheme
	lt  := defaultLabelTheme
	bt  := defaultButtonTheme
	tit := defaultTextInputTheme
	
	dt := jsonTheme{
		Map        : defaultMapTheme,
		InfoPanel  : pt,
		CtrlPanel  : pt,
		Generation : lt,
		Speed      : lt,
		Zoom       : lt,
		PlayToggle : bt,
		SlowDown   : bt,
		SpeedUp    : bt,
		ResetState : bt,
		RandomState: bt,
		FF_I       : bt,
		FF_X       : bt,
		FF_L       : bt,
		FF_C       : bt,
		FF_M       : bt,
		Rules      : tit,
	}

	return SaveTheme(&dt, path)
}

/* Saves the JSON data structure at the specified path. Errors are related to JSON data corruption or file handling. */
func SaveTheme(jt *jsonTheme, path string) error {
	stream, err := json.MarshalIndent(jt, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, stream, 0644)
}
