package main

import "testing"

/* Tests whether the number closest to mean is properly picked. */
func TestGetClosestToMean(t *testing.T) {
	type testCase struct {
		input    []float32
		expected float32
	}

	cases := []testCase{
		{
			input:    []float32{5, 4, 3, 7},
			expected: 5,
		},
		{
			input:    []float32{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60},
			expected: 15,
		},
		{
			input:    []float32{},
			expected: 0,
		},
		{
			input:    []float32{42},
			expected: 42,
		},
		{
			input:    []float32{2, -17, 5, 3, 5},
			expected: 2,
		},
	}

	for _, c := range cases {
		out := GetClosestToMean(c.input)
		if out != c.expected {
			t.Fatalf("Failed for %v: got %v expected %v\n", c.input, out, c.expected)
		}
	}
}

/* Tests whether GetCommonDivisors returns appropriate common factors. */
func TestGetCommonDivisors(t *testing.T) {
	type testCase struct {
		input    []float32
		expected []float32
		min, max float32
	}

	cases := []testCase{
		{
			input:    []float32{60},
			expected: []float32{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60},
			min:      1,
			max:      60,
		},
		{
			input:    []float32{480, 640},
			expected: []float32{2, 4, 5, 8, 10, 16, 20},
			min:      2,
			max:      20,
		},
		{
			input:    []float32{360, 480, 640, 720, 1080, 1280, 1920},
			expected: []float32{2, 4, 5, 8, 10, 20},
			min:      2,
			max:      20,
		},
	}

	failMsg := "Failed for %v: got %v expected %v\n"

	for _, c := range cases {
		divs := GetCommonDivisors(c.min, c.max, c.input...)

		if len(divs) != len(c.expected) {
			t.Fatalf(failMsg, c.input, divs, c.expected)
		}

		for i := range divs {
			if divs[i] < c.min || divs[i] > c.max {
				t.Fatalf(failMsg+"Range error", c.input, divs, c.expected)
			}

			if divs[i] != c.expected[i] {
				t.Fatalf(failMsg, c.input, divs, c.expected)
			}
		}
	}
}
