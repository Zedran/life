package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Holds the short commit of the build
var Version string

func main() {
	fmt.Println(Version)

	g := NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
