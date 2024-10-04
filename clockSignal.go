package main

// Indicates what action should be taken in regards to the current Clock.Ticks value
type ClockSignal uint8

const (
	// Do nothing
	WAIT ClockSignal = iota

	// Trigger an event now
	TRIGGER
)
