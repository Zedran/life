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
		"23/3"  : {[]uint8{2, 3}, []uint8{3}}, // Typical rules
		"23/"   : {[]uint8{2, 3}, []uint8{ }}, // No dead rules
		"/3"    : {[]uint8{    }, []uint8{3}}, // No alive rules
		"233/33": {[]uint8{2, 3}, []uint8{3}}, // Duplicates
		"/"     : {[]uint8{    }, []uint8{ }}, // No rules
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
				break
			}
		}

		for i := range r.Die {
			if r.Die[i] != out.Die[i] {
				badContent = true
				break
			}
		}

		if badContent {
			t.Errorf("Failed for %s.\n  Got: %v / %v\n  Expected: %v / %v", gc, r.Live, r.Die, out.Live, out.Die)
		}
	}
}
