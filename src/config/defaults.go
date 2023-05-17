package config

var (
	defaultConfig = jsonConfig{
		WorldSize: "auto",
		Language : "",
		Theme    : "",
		Window   : Window{
			W: 720,
			H: 480,
		},
	}
)
