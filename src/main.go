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

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
