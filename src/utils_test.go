package main

import "testing"

/* Tests whether GetCommonDivisors returns appropriate common factors. */
func TestGetCommonDivisors(t *testing.T) {
	type testCase struct {
		input    []float32
		expected []float32
	}

	var (
		min float32 =  2
		max float32 = 20
	)

	cases := []testCase{
		{
			input   : []float32{480, 640}, 
			expected: []float32{2, 4, 8, 10, 16, 20},
		},
		{
			input   : []float32{360, 480, 640, 720, 1080, 1280, 1920}, 
			expected: []float32{2, 4, 8, 10, 20},
		},
	}

	failMsg := "Failed for %v: got %v expected %v\n"

	for _, c := range cases {
		divs := GetCommonDivisors(min, max, c.input...)

		if len(divs) != len(c.expected) {
			t.Fatalf(failMsg, c.input, divs, c.expected)
		}

		for i := range divs {
			if divs[i] < min || divs[i] > max {
				t.Fatalf(failMsg + "Range error", c.input, divs, c.expected)
			}

			if divs[i] != c.expected[i] {
				t.Fatalf(failMsg, c.input, divs, c.expected)
			}
		}
	}
}
