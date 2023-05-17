package config

var (
	defaultConfig = jsonConfig{
		WorldSize: "auto",
		Language : "en",
		Theme    : "default",
		Window   : Window{
			W: 720,
			H: 480,
		},
	}
)
