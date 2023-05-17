package lang

import (
	"encoding/json"
	"os"
)

/* Language of the game. */
type Language struct {
	// Window title
	Title      string `json:"title"`

	// Generation LabeledDisplay
	Generation string `json:"generation"`

	// Speed LabeledDisplay
	Speed      string `json:"speed"`

	// Zoom LabeledDisplay
	Zoom       string `json:"zoom"`

	Rules      string `json:"rules"`
}

/* Loads language from file. If the file does not exist, returns default language. */
func LoadLanguage(path string) *Language {
	stream, err := os.ReadFile(path)
	if err != nil {
		return &defaultLanguage
	}

	var lang Language

	if err = json.Unmarshal(stream, &lang); err != nil {
		return &defaultLanguage
	}

	return &lang
}

/* Saves the unexported default language data. */
func SaveDefault(path string) error {
	return SaveLanguage(&defaultLanguage, path)
}

/* Saves the JSON data structure to specified path. Errors are related to JSON data corruption or file handling. */
func SaveLanguage(l *Language, path string) error {
	stream, err := json.MarshalIndent(l, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, stream, 0644)
}
