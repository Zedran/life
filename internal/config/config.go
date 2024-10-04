package config

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Zedran/life/internal/config/lang"
	"github.com/Zedran/life/internal/config/theme"
)

const (
	// Top of the config file path
	CONFIG_DIR string = "config"

	// Languages directory
	LANG_DIR string = CONFIG_DIR + "/languages"

	// Theme directory
	THEME_DIR string = CONFIG_DIR + "/themes"

	// Config file
	CONFIG_PATH string = CONFIG_DIR + "/config.json"

	// Default language file
	DEFAULT_LANG_PATH string = LANG_DIR + "/en.json"

	// Default theme file
	DEFAULT_THEME_PATH string = THEME_DIR + "/default.json"

	// Default minimum zoom value
	DEFAULT_ZOOM_MIN float32 = 4

	// Default maximum zoom value
	DEFAULT_ZOOM_MAX float32 = 20

	// Lowest zoom value allowed
	ZOOM_MIN_LIMIT float32 = 1
)

/* Configuration of the game. */
type Config struct {
	// Number of cells in a row
	WorldSize int

	// Minimum zoom value for the map, limited to ZOOM_MIN_LIMIT
	ZoomMin float32

	// Maximum zoom value for the map, unlimited in code, but the cap is always
	// the greatest common factor of Window.W and Window.H, regardless of the higher setting
	ZoomMax float32

	// Language of the game
	Language *lang.Language

	// Color theme of the game
	Theme *theme.Theme

	// Window configuration
	Window *Window
}

/* Builds the config directory tree. */
func createDirTree(root string) error {
	const perm fs.FileMode = 0755

	if err := os.MkdirAll(filepath.Join(root, LANG_DIR), perm); err != nil {
		return err
	}

	if err := os.Mkdir(filepath.Join(root, THEME_DIR), perm); err != nil {
		return err
	}

	return nil
}

/* Returns true if config directory tree exists and is complete. */
func dirTreeExists(root string) bool {
	for _, d := range []string{LANG_DIR, THEME_DIR} {
		if _, err := os.Stat(filepath.Join(root, d)); err != nil {
			return false
		}
	}

	return true
}

/* Loads the config from file. If the file does not exist, returns default config. */
func LoadConfig() *Config {
	root, err := GetRootDir()
	if err != nil {
		return defaultConfig.ToConfig(root)
	}

	var jc jsonConfig

	stream, err := os.ReadFile(filepath.Join(root, CONFIG_PATH))
	if err != nil {
		return defaultConfig.ToConfig(root)
	}

	if err = json.Unmarshal(stream, &jc); err != nil {
		return defaultConfig.ToConfig(root)
	}

	return jc.ToConfig(root)
}

/* Saves the unexported default config data. */
func SaveDefault(path string) error {
	return SaveConfig(&defaultConfig, path)
}

/* Saves the JSON data structure at the specified path. Errors are related to JSON data corruption or file handling. */
func SaveConfig(jc *jsonConfig, path string) error {
	stream, err := json.MarshalIndent(jc, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, stream, 0644)
}

/* Builds the config directory tree and writes all the default files into it. Files that already exist are not overwritten. */
func WriteDefaults() error {
	root, err := GetRootDir()
	if err != nil {
		return err
	}

	if !dirTreeExists(root) {
		if err := createDirTree(root); err != nil {
			return err
		}
	}

	resources := map[string]func(string) error{
		filepath.Join(root, CONFIG_PATH):        SaveDefault,
		filepath.Join(root, DEFAULT_LANG_PATH):  lang.SaveDefault,
		filepath.Join(root, DEFAULT_THEME_PATH): theme.SaveDefault,
	}

	for path, saveFunc := range resources {
		if _, err := os.Stat(path); err != nil {
			if err := saveFunc(path); err != nil {
				return err
			}
		}
	}

	return nil
}
