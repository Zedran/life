package world

import "sync"

/*
BufferPool provides World with a set of three reusable buffers,
eliminating the need to allocate new buffer every generation.
*/
type BufferPool struct {
	// Holds current state of the World
	state []State

	// During world.Update, new state of the World is written into BufferPool.buffer
	buffer []State

	// Unused buffer, cleared during World.Update
	spare []State
}

/*
Clears currently unused buffer. This method uses World.wg (sync.WaitGroup),
which allows bp.spare to be wiped in parallel to the loop in World.Update.
*/
func (bp *BufferPool) ClearSpareBuffer(wg *sync.WaitGroup) {
	for i := range bp.spare {
		bp.spare[i] = DEAD
	}
	wg.Done()
}

/* Returns a pointer to the current buffer. */
func (bp *BufferPool) GetCurrentBuffer() []State {
	return bp.buffer
}

/* Returns a pointer to the current state of the World. */
func (bp *BufferPool) GetCurrentState() []State {
	return bp.state
}

/* Returns a pointer to the currently unused buffer. */
func (bp *BufferPool) GetSpareBuffer() []State {
	return bp.spare
}

/*
Switches the buffer pointers in the following way:
  - buffer becomes new state
  - spare is now buffer
  - state becomes spare
*/
func (bp *BufferPool) NextState() {
	pPrevState := bp.state

	bp.state = bp.buffer
	bp.buffer = bp.spare

	bp.spare = pPrevState
}

/* Creates new BufferPool given buffer size (size of the world). */
func NewBufferPool(bufferSize int) *BufferPool {
	return &BufferPool{
		state:  newStateBuffer(bufferSize),
		buffer: newStateBuffer(bufferSize),
		spare:  newStateBuffer(bufferSize),
	}
}
