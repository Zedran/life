package config

import (
	"path/filepath"
	"strconv"

	"github.com/Zedran/life/src/config/lang"
	"github.com/Zedran/life/src/config/theme"
)

/* A JSON representation of Config. */
type jsonConfig struct {
	WorldSize string `json:"world_size"`

	ZoomMin   float32 `json:"zoom_min"`
	ZoomMax   float32 `json:"zoom_max"`
	
	Language  string `json:"language"`
	Theme     string `json:"theme"`

	Window    Window `json:"window"`

}

/*
	Converts JSON struct into Config. The key points of this process are as follows:

	WorldSize cannot be less than half of the Window's pixel width. If "auto" keyword is specified 
	or if the WorldSize does not meet the above requirements, it is set to default 
	(a cell count equal to half of Window.W or Window.H [px] - whichever is greater).

	If language file does not exist or its structure is invalid, a default one is loaded.

	If theme file does not exist or its structure is invalid, a default one is loaded.

	If ZOOM_MIN value is lower than ZOOM_MIN_LIMIT, its value will be changed to lowest within limit.
	Setting ZOOM_MIN value to ZOOM_MIN_LIMIT disables theme.MapTheme.Border.

	If ZOOM_MAX value is lower than ZOOM_MIN, its value will be changed to ZOOM_MIN.
	Zoom does not have a hardcoded upper limit, but its practical max is the greatest common
	divisor of window's width and height [px].
*/
func (jc *jsonConfig) ToConfig(rootDir string) *Config {
	if jc.ZoomMin < ZOOM_MIN_LIMIT {
		jc.ZoomMin = ZOOM_MIN_LIMIT
	}

	if jc.ZoomMax < jc.ZoomMin {
		jc.ZoomMax = jc.ZoomMin
	}

	n, err := strconv.Atoi(jc.WorldSize)

	if err != nil || jc.WorldSize == "auto" || n < GetDefaultWorldSize(&jc.Window, jc.ZoomMin) {
		n = GetDefaultWorldSize(&jc.Window, jc.ZoomMin)
	}
	
	lang  := lang.LoadLanguage(filepath.Join(rootDir, LANG_DIR,  VerifyCfgFileExt(jc.Language)))
	theme := theme.LoadTheme(  filepath.Join(rootDir, THEME_DIR, VerifyCfgFileExt(jc.Theme   )))

	if jc.ZoomMin == ZOOM_MIN_LIMIT {
		theme.MapTheme.Border = false
	}

	return &Config{
		WorldSize: n,
		ZoomMin  : jc.ZoomMin,
		ZoomMax  : jc.ZoomMax,
		Language : lang,
		Theme    : theme,
		Window   : &jc.Window,
	}
}
