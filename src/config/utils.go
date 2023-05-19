package config

import (
	"os"
	"path/filepath"
)

/*
	Returns the default world size. The default size is twice the amount of cells that fits into greater window dimension
	at minimum zoom level. In other words, the length of Map's side (Map.Size - cell count) equals to half of the Window.W
	or Window.H [px] - whichever is greater.
*/
func GetDefaultWorldSize(w *Window, zoomMin float32) int {
	var dim float32

	if w.W > w.H {
		dim = w.W
	} else {
		dim = w.H
	}

	return int(2 * dim / zoomMin)
}

/* Returns the root directory in which the executable is located. */
func GetRootDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exePath), nil
}

/* Verifies that the path ends with '.json' extension and appends it if not found. */
func VerifyCfgFileExt(path string) string {
	const configFileExt = ".json"

	if filepath.Ext(path) != configFileExt {
		return path + configFileExt
	}

	return path
}
