package world

import (
	"math/rand"
	"sync"
)

/*
Size of the cell map border that is not checked during World.Update.
Skipping at least one row and column prevents the out of bounds
access to the slice, so there is no need to check for it.
*/
const PADDING int = 1

/* Represents the world of the game - a square grid of cells. */
type World struct {
	// BufferPool for the State array
	bp *BufferPool

	// The number of cells in one row / column
	Size int

	// Current generation of the world
	Generation uint64

	// State of the cells, points to World.bp.state
	Cells []State

	// Game rules currently in effect
	Rules *Rules

	// WaitGroup for World.Update method
	wg *sync.WaitGroup

	// Indicates whether the World is currently updating (World.wg waits)
	Working bool
}

/*
Randomly fills the world with life.
The accepted values of density parameter fall in range <1; 9>.
The lower the density, the more scarce life becomes.
*/
func (w *World) RandomState(density int) {
	const DENSITY_MAX int = 10

	if density < 1 {
		density = 1
	} else if density >= DENSITY_MAX {
		density = DENSITY_MAX - 1
	}

	w.Reset()

	for i := range w.Cells {
		if rand.Intn(DENSITY_MAX-density) == 0 {
			w.Cells[i] = State(rand.Intn(2))
		}
	}
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
	w.Working = true

	w.Generation++

	buffer := w.bp.GetCurrentBuffer()

	w.wg.Add(1 + w.Size - PADDING*2)

	go w.bp.ClearSpareBuffer(w.wg)

	for y := w.Size * PADDING; y < len(w.Cells)-w.Size*PADDING; y += w.Size {

		go func(rowStart int) {

			for i := rowStart + PADDING; i < rowStart+w.Size-PADDING; i++ {
				// Neighbour count
				nc :=
					w.Cells[i-1-w.Size] +
						w.Cells[i-w.Size] +
						w.Cells[i+1-w.Size] +
						w.Cells[i-1] +
						w.Cells[i+1] +
						w.Cells[i-1+w.Size] +
						w.Cells[i+w.Size] +
						w.Cells[i+1+w.Size]

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
			w.wg.Done()
		}(y)
	}
	w.wg.Wait()

	w.bp.NextState()

	w.Cells = w.bp.GetCurrentState()

	w.Working = false
}

/* Progresses the World by n generations. If stop signal is received from the channel, jump task is dropped. */
func (w *World) UpdateBy(stop *chan bool, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-*stop:
			return
		default:
			w.Update()
		}
	}
}

/* Creates new world of specified size. */
func Genesis(worldSize int) *World {
	cells := make([]State, worldSize*worldSize, worldSize*worldSize)

	rules, _ := NewRules(DEFAULT_RULES)

	return &World{
		bp:         NewBufferPool(len(cells)),
		Generation: 0,
		Size:       worldSize,
		Cells:      cells,
		Rules:      rules,
		wg:         &sync.WaitGroup{},
		Working:    false,
	}
}
