package world

import "sync"

/* Size of the cell map border that is not checked during World.Update.
   Skipping at least one row and column prevents the out of bounds
   access to the slice, so there is no need to check for it.
*/
const PADDING int = 1

/* Represents the world of the game - a square grid of cells. */
type World struct {
	// The number of cells in one row / column
	Size       int

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
		rowLen = int(w.Size)
		buffer = make([]State, len(w.Cells), len(w.Cells))
	)

	var wg sync.WaitGroup

	wg.Add(w.Size - PADDING * 2)

	for y := rowLen * PADDING; y < len(w.Cells) - rowLen * PADDING; y += rowLen  {
		
		go func(rowStart int) {

			for i := rowStart + PADDING; i < rowStart + rowLen - PADDING; i++ {
				// Neighbour count
				nc := 
					w.Cells[i - 1 - rowLen] + 
					w.Cells[i     - rowLen] + 
					w.Cells[i + 1 - rowLen] + 
					w.Cells[i - 1         ] + 
					w.Cells[i + 1         ] + 
					w.Cells[i - 1 + rowLen] + 
					w.Cells[i     + rowLen] + 
					w.Cells[i + 1 + rowLen]

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
							buffer[i] = ALIVE
							break
						}
					}
				}
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
	w.Cells = buffer
}

/* Creates new world of specified size. */
func Genesis(worldSize int) *World {
	cells := make([]State, worldSize * worldSize, worldSize * worldSize)

	rules, _ := NewRules(DEFAULT_RULES)

	return &World{
		Generation: 0,
		Size      : worldSize,
		Cells     : cells,
		Rules     : *rules,
	}
}
