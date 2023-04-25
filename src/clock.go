package main

import "github.com/hajimehoshi/ebiten/v2"

/* Clock limits the event occurence rate to the specified number per second. */
type Clock struct {
	// Index on the SpeedDial, pointing to the currently set speed
	CurrentSpeed  int

	// A collection of possible speeds expressed as game update ticks between the events
	SpeedDial     []int

	// Ticks since last event
	Ticks         int
}

/*
	Adjusts the speed, moving up or down the Clock.SpeedDial slice. Does nothing if the value
	of direction is 0 or puts Clock.CurrentSpeed beyond the bounds of the slice.
*/
func (c *Clock) AdjustSpeed(direction int) {
	direction = -(direction)

	if (direction == 0) || (c.CurrentSpeed + direction < 0 || c.CurrentSpeed + direction >= len(c.SpeedDial)) {
		return
	}

	c.CurrentSpeed += direction
}

/* Returns the current speed expressed as events per second. */
func (c *Clock) GetEventsPerSec() int {
	return ebiten.TPS() / c.SpeedDial[c.CurrentSpeed]
}

/*
   Increments Clock.Ticks or, if the appropriate number of ticks have passed, resets the counter
   and returns the TRIGGER ClockSignal.
*/
func (c *Clock) Tick() ClockSignal {
	if c.Ticks == c.SpeedDial[c.CurrentSpeed] - 1 {
		c.Ticks = 0
		return TRIGGER
	}

	c.Ticks++
	return WAIT
}

/* Creates new Clock. The possible speed values are dependent on game TPS (only integer factors are used). */
func NewClock() *Clock {
	var c Clock

	dial := GetCommonDivisors(1, float32(ebiten.TPS()), float32(ebiten.TPS()))

	c.SpeedDial = make([]int, len(dial))

	for i := range dial {
		c.SpeedDial[i] = int(dial[i])
	}

	c.CurrentSpeed = Index(dial, GetClosestToMean(dial))

	return &c
}
