package ui

/* Represents the state of the Controller. It is emitted outwards on Controller.Emit call. */
type UIResponse struct {
	// Current rules read from the TextInput
	Rules string

	// Signal passed since last Controller.Emit call
	Signal UISignal
}
