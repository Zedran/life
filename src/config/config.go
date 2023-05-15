package config

import (
	"path/filepath"

	"github.com/Zedran/life/src/config/lang"
	"github.com/Zedran/life/src/config/theme"
)

const (
	// Top of the config file path
	CONFIG_DIR      string  = "config"

	// Config file
	CONFIG_PATH     string  = CONFIG_DIR + "/config.json"

	// Languages directory
	LANG_DIR        string  = CONFIG_DIR + "/languages"

	// Theme directory
	THEME_DIR       string  = CONFIG_DIR + "/themes"

	// Minimum zoom value
	ZOOM_MIN        float32 =  4

	// Maximum allowed zoom
	ZOOM_MAX        float32 = 20
)

/* Configuration of the game. */
type Config struct {
	// Number of cells in a row
	WorldSize int

	// Language of the game
	Language  *lang.Language

	// Color theme of the game
	Theme     *theme.Theme

	// Window configuration
	Window    *Window
}

/* Loads the config from file. If the file does not exist, returns default config. */
func LoadConfig() *Config {
	return LoadDefaultConfig()
}

/* Returns the default game configuration. */
func LoadDefaultConfig() *Config {
	return &Config{
		WorldSize: 720 / int(ZOOM_MIN / 2),
		Language : lang.LoadLanguage(filepath.Join(LANG_DIR, "")),
		Theme    : theme.LoadTheme(filepath.Join(THEME_DIR, "")),
		Window   : &Window{
			W    : 720,
			H    : 480,
		},
	}
}
