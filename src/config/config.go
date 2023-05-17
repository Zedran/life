package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Zedran/life/src/config/lang"
	"github.com/Zedran/life/src/config/theme"
)

const (
	// Top of the config file path
	CONFIG_DIR  string  = "config"

	// Config file
	CONFIG_PATH string  = CONFIG_DIR + "/config.json"

	// Languages directory
	LANG_DIR    string  = CONFIG_DIR + "/languages"

	// Theme directory
	THEME_DIR   string  = CONFIG_DIR + "/themes"

	// Minimum zoom value
	ZOOM_MIN    float32 =  4

	// Maximum allowed zoom
	ZOOM_MAX    float32 = 20
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
	root, err := GetRootDir()
	if err != nil {
		return defaultConfig.ToConfig(root)
	}

	var	jc jsonConfig

	stream, err := os.ReadFile(filepath.Join(root, CONFIG_PATH))
	if err != nil {
		return defaultConfig.ToConfig(root)
	}

	if err = json.Unmarshal(stream, &jc); err != nil {
		return defaultConfig.ToConfig(root)
	}

	return jc.ToConfig(root)
}
