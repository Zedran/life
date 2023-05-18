package config

import (
	"os"
	"path/filepath"
)

/*
	Returns the default world size. The default size is twice the amount of cells that fits into windowWidth
	at minimum zoom level. In other words, the length of Map's side (Map.Size) equals to half of the Window.W [px].
*/
func GetDefaultWorldSize(windowWidth, zoomMin float32) int {
	return int(2 * windowWidth / zoomMin)
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
