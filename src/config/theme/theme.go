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

/*
	Saves the JSON data structure at the specified path. Errors are related to JSON data corruption
	or file handling.
*/
func SaveTheme(t *jsonTheme, path string) error {
	stream, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, stream, 0644)
}
