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
	
	Language  string `json:"language"`
	Theme     string `json:"theme"`

	Window    Window `json:"window"`

}

/*
	Converts JSON struct into Config. The key points of this process are as follows:

	WorldSize cannot be less than half of the Window's pixel width. If "auto" keyword is specified 
	or if the WorldSize does not meet the above requirements, it is set to default 
	(a cell count equal to a half of Window.W [px].)

	If language file does not exist or its structure is invalid, a default one is loaded.

	If theme file does not exist or its structure is invalid, a default one is loaded.
*/
func (jc *jsonConfig) ToConfig(rootDir string) *Config {
	n, err := strconv.Atoi(jc.WorldSize)

	if err != nil || jc.WorldSize == "auto" || n < GetDefaultWorldSize(jc.Window.W) {
		n = GetDefaultWorldSize(jc.Window.W)
	}
	
	lang  := lang.LoadLanguage(filepath.Join(rootDir, LANG_DIR,  VerifyCfgFileExt(jc.Language)))
	theme := theme.LoadTheme(  filepath.Join(rootDir, THEME_DIR, VerifyCfgFileExt(jc.Theme   )))

	return &Config{
		WorldSize: n,
		Language : lang,
		Theme    : theme,
		Window   : &jc.Window,
	}
}
