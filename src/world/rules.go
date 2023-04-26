package world

import (
	"errors"
	"strconv"
	"strings"
)

// Default rules of Conway's Game of Life
const DEFAULT_RULES = "23/3"

// Error returned if specified rules do not have correct format
var errInvalidRules = errors.New("invalid rules")

/*
   Rules of the game. They specify a number of neighbours required 
   for the cell to live or die on transition to next generation.
*/
type Rules struct {
	// How many neighbours ensure survival of the cell
	Live []uint8

	// How many neighbours are required for the cell to die
	Die  []uint8
}

/*
   Creates new Rules struct. Error returned means that specified string
   is not formatted correctly, i.e. <live>/<die>. '/' must always be present.
*/
func NewRules(ruleString string) (*Rules, error) {
	const RULE_SEP string = "/"

	if strings.Count(ruleString, RULE_SEP) != 1 {
		return nil, errInvalidRules
	}

	ruleString = strings.Replace(ruleString, " ", "", -1)

	vals := strings.Split(ruleString, RULE_SEP)

	var (
		err error
		r   Rules
	)

	if r.Live, err = parseRuleSubstring(vals[0]); err != nil {
		return nil, errInvalidRules
	}

	if r.Die, err = parseRuleSubstring(vals[1]); err != nil {
		return nil, errInvalidRules
	}

	return &r, nil
}

/*
   Makes an in depth validation of rules substring (either 'live' or 'die' part)
   and returns a slice with number of neighbours required to trigger the event.
*/
func parseRuleSubstring(sub string) ([]uint8, error) {
	s := make([]uint8, 0)

	// Empty rules are possible
	if len(sub) == 0 {
		return s, nil
	}

	for _, v := range strings.Split(sub, "") {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		// The number of neighbours is within range <0; 8>
		if n < 0 || n > 8 {
			return nil, errInvalidRules
		}

		// Check for duplicates
		if !Contains(s, uint8(n)) {
			s = append(s, uint8(n))
		}
	}

	return s, nil
}
