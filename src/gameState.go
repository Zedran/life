package main

// Indicates the current state of the Game
type GameState uint8

const (
	// The game is running (world.World is progressing)
	RUN   GameState = iota

	// The game is paused (world.World is not progressing)
	PAUSE
)
