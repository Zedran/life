package config

const (
	CONFIG_LOCATION string  = "config.json"

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
	Language  *Language

	// Color theme
	Theme     *Theme

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
		Language : &Language{
			Title: "Game of Life",
		},
		Theme    : LoadDefaultTheme(),
		Window   : &Window{
			W    : 720,
			H    : 480,
		},
	}
}
