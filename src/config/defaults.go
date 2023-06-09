package config

var (
	defaultConfig = jsonConfig{
		WorldSize: "auto",
		ZoomMin  : DEFAULT_ZOOM_MIN,
		ZoomMax  : DEFAULT_ZOOM_MAX,
		Language : "en",
		Theme    : "default",
		Window   : Window{
			W: 720,
			H: 480,
		},
	}
)
