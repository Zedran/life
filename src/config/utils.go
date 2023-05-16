package config

import (
	"os"
	"path/filepath"
)

/*
	Returns the default world size. The default size is twice the amount of cells that fits into windowWidth
	at minimum zoom level. In other words, the length of Map's side (Map.Size) equals to half of the Window.W [px].
*/
func GetDefaultWorldSize(windowWidth float32) int {
	return int(windowWidth / (ZOOM_MIN / 2))
}

/* Returns the root directory in which the executable is located. */
func GetRootDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exePath), nil
}