package world

/* The size of the cell map border that is omitted during World.Update.
   Skipping at least one row eliminates the risk of out of bounds access
   to the array. There is no need to waste time checking for it.
*/
const PADDING int = 1

/* Represents the world of the game - a square grid of cells. */
type World struct {
	// The number of cells in one row / column
	Size       uint64

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

/* Makes the world transition to next generation, updating the cell array. */
func (w *World) Update() {
	w.Generation++

	var (
		pitch  = int(w.Size)
		buffer = make([]State, len(w.Cells), len(w.Cells))
	)

	for i := pitch + PADDING; i <= int(w.Size) - pitch - PADDING; i += PADDING * 2 + 1  {
		
			rowStart := i

			for ; i < rowStart + pitch - PADDING * 2 - 1; i++ {
				// Neighbour count
				nc := 
					w.Cells[i - 1 - pitch] + 
					w.Cells[i     - pitch] + 
					w.Cells[i + 1 - pitch] + 
					w.Cells[i - 1        ] + 
					w.Cells[i + 1        ] + 
					w.Cells[i - 1 + pitch] + 
					w.Cells[i     + pitch] + 
					w.Cells[i + 1 + pitch]

				switch w.Cells[i] {
				case ALIVE:
					for _, lr := range w.Rules.Live {
						if State(lr) == nc {
							buffer[i] = ALIVE
							break
						}
					}
				case DEAD:
					for _, dr := range w.Rules.Die {
						if State(dr) == nc {
							buffer[i] = DEAD
							break
						}
					}
				}
			}
	}
	w.Cells = buffer
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
