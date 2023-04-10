package world

// Indicates the state of a cell
type State uint8

const (
	DEAD  = State(iota)
	ALIVE
)
