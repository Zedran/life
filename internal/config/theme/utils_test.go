package theme

import "testing"

/* Tests the color code parsing function. */
func TestDecodeColors(t *testing.T) {
	badCases := []string{
		"",           // Empty
		"12",         // Too short
		"ghi",        // Invalid characters
		"0123456789", // Too long
	}

	for _, bc := range badCases {
		col, err := decodeColor(bc)
		if err == nil {
			cval := []byte{col.R, col.G, col.B, col.A}

			t.Errorf("Did not fail on '%s'. Got %v\n", bc, cval)
		}
	}

	cases := map[string][]byte{
		"deab12":    {0xde, 0xab, 0x12, 0xff}, // RGB
		"#deab12":   {0xde, 0xab, 0x12, 0xff}, // RGB with prefix
		"12345678":  {0x12, 0x34, 0x56, 0x78}, // RGBA
		"#12345678": {0x12, 0x34, 0x56, 0x78}, // RGBA with prefix
		"FFFFFF":    {0xff, 0xff, 0xff, 0xff}, // Upper case
	}

	for c, eo := range cases {
		col, err := decodeColor(c)
		if err != nil {
			t.Errorf("Failed on '%s'. Error: %s\n", c, err.Error())
		}

		cval := []byte{col.R, col.G, col.B, col.A}

		for i := range cval {
			if cval[i] != eo[i] {
				t.Errorf("Failed on '%s'. Expected %v, got %v.\n", c, eo, cval)
			}
		}
	}
}
