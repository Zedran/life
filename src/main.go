package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGame()

	g.World.RandomState(5)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
