package ui

/*
	Signal received by Controller. Each signal corresponds with specific widget
	and allows the Game to determine in what way its state should be modified.
*/
type UISignal uint8

const (
	// No signal is currently stored (UI was not interacted with)
	NONE         UISignal = iota

	// Toggle Play / Pause
	PLAY_TOGGLE

	// Decrease speed
	SLOW_DOWN
	
	// Increase speed
	SPEED_UP

	// Reset the world (set the state of all cells to DEAD)
	RESET_STATE

	// Set random state
	RANDOM_STATE

	// Fast forward 1 generation
	FF_I

	// Fast forward 10 generations
	FF_X

	// Fast forward 50 generations
	FF_L

	// Fast forward 100 generations
	FF_C

	// Fast forward 1000 generations
	FF_M

	// Indicates that new rules were requested
	NEW_RULES
)
