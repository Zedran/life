package theme

import (
	"encoding/hex"
	"errors"
	"image/color"
	"strings"
)

// Returned from decodeColor if color code is not properly structured
var errMalformedColorCode = errors.New("malformed color code")

/*
	Converts a color code to color.RGBA struct. Accepts the following code formats: 
		#0a0b0c   - RGB
		#0a0b0cff - RGBA
	Ignores case. Hash prefix is optional.
*/
func decodeColor(code string) (*color.RGBA, error) {
	code = strings.Replace(code, "#", "", -1)

	x, err := hex.DecodeString(code)
	if err != nil {
		return nil, err
	}

	if len(x) == 3 {
		x = append(x, 0xff)
	} else if len(x) != 4 {
		return nil, errMalformedColorCode
	}

	return &color.RGBA{x[0], x[1], x[2], x[3]}, err
}

/* Converts color.RGBA to color.NRGBA. */
func rgbaToNRGBA(c *color.RGBA) *color.NRGBA {
	nrgba := (color.NRGBA)(*c)
	
	return &nrgba
}
