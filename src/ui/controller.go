package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

/* Controller for the UI. */
type Controller struct {
	// A widget which registered most recent left mouse button down
	lmbDownOn *widget.Widget

	// Current rules read from the TextInput
	rules string

	// Signal passed since last Controller.Emit call
	signal UISignal
}

/* Sets c.lmbDownOn to nil. */
func (c *Controller) ClearLMBPointer() {
	c.lmbDownOn = nil
}

/* Clears the signal. Called on Controller.Emit*/
func (c *Controller) clearSignal() {
	c.signal = NONE
}

/* Emits the current UI state outwards in form of UIResponse. */
func (c *Controller) Emit() *UIResponse {
	defer c.clearSignal()

	return &UIResponse{
		Rules:  c.rules,
		Signal: c.signal,
	}
}

func (c *Controller) GetRules() string {
	return c.rules
}

func (c *Controller) GetSignal() UISignal {
	return c.signal
}

/* Returns true if w and c.lmbDownOn point to the same widget. */
func (c *Controller) IsLMBDownOn(w *widget.Widget) bool {
	return w == c.lmbDownOn
}

/* Stores the pointer to w in c.lmbDownOn. */
func (c *Controller) SetLMBDownOn(w *widget.Widget) {
	c.lmbDownOn = w
}

func (c *Controller) SetRules(rules string) {
	c.rules = rules
}

func (c *Controller) SetSignal(s UISignal) {
	c.signal = s
}

/* Creates a new Controller structure. */
func NewController(rules string) *Controller {
	return &Controller{
		lmbDownOn: nil,
		rules:     rules,
		signal:    NONE,
	}
}
