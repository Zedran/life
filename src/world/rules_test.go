package world

import "testing"

/* Tests whether the rules in string form are correctly parsed into Rules struct. */
func TestNewRules(t *testing.T) {
	badCases := []string{"1234567", "12/34/56", "abc", "12avc", "1234/56.", "AB", "123A/4", "23/9"}

	for _, bc := range badCases {
		if r, err := NewRules(bc); err == nil {
			t.Errorf("Did not fail for %s. Got %v.", bc, r)
		}
	}

	goodCases := map[string]Rules{
		"23/3"  : {[]uint8{2, 3}, []uint8{3}, "23/3"  }, // Typical rules
		"233/33": {[]uint8{2, 3}, []uint8{3}, "233/33"}, // Duplicates
		"23/"   : {[]uint8{2, 3}, []uint8{ }, "23/"   }, // No dead rules
		"/3"    : {[]uint8{    }, []uint8{3}, "/3"    }, // No alive rules
		"/"     : {[]uint8{    }, []uint8{ }, "/"     }, // No rules
		" 1 / 1": {[]uint8{1   }, []uint8{1}, " 1 / 1"}, // Spaces
	}

	for gc, out := range goodCases {
		r, err := NewRules(gc)
		if err != nil {
			t.Errorf("Failed for %s. Error: %s", gc, err.Error())
		}

		badContent := false
		
		if len(r.Live) != len(out.Live) || len(r.Die) != len(out.Die) {
			badContent = true
		}

		for i := range r.Live {
			if r.Live[i] != out.Live[i] {
				badContent = true
			}
		}

		for i := range r.Die {
			if r.Die[i] != out.Die[i] {
				badContent = true
			}
		}

		if badContent {
			t.Errorf("Failed for %s.\n  Got: %v / %v\n  Expected: %v / %v", gc, r.Live, r.Die, out.Live, out.Die)
		}
	}
}
