package world

/* Represents the world of the game - a square grid of cells. */
type World struct {
	// The number of cells in one row / column
	Size  uint64

	// Current generation of the world
	Generation uint64

	// State of the cells
	Cells      []State

	// Game rules currently in effect
	Rules      Rules
}

/* Reset the world by clearing the cell states and setting the generation number to zero. */
func (w *World) Reset() {
	w.Generation = 0

	for i := range w.Cells {
		w.Cells[i] = 0
	}
}

/* Creates new world of specified size. */
func Genesis(worldSize uint64) *World {
	cells := make([]State, worldSize * worldSize, worldSize * worldSize)

	rules, _ := NewRules(DEFAULT_RULES)

	return &World{
		Generation: 0,
		Size      : worldSize,
		Cells     : cells,
		Rules     : *rules,
	}
}
