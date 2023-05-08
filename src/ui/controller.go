package ui

/* Controller for the UI. */
type Controller struct {
	// Current rules read from the TextInput
	rules  string

	// Signal passed since last Controller.Emit call
	signal UISignal
}

/* Clears the signal. Called on Controller.Emit*/
func (c *Controller) clearSignal() {
	c.signal = NONE
}

/* Emits the current UI state outwards in form of UIResponse. */
func (c *Controller) Emit() *UIResponse {
	defer c.clearSignal()

	return &UIResponse{
		Rules : c.rules,
		Signal: c.signal,
	}
}

func (c *Controller) GetRules() string {
	return c.rules
}

func (c *Controller) GetSignal() UISignal {
	return c.signal
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
		rules : rules,
		signal: NONE,
	}
}
