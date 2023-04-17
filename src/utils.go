package main

import "math"

/* Returns a slice containing common divisors of numbers in nums between min and max. */
func GetCommonDivisors(min, max float32, nums ...float32) []float32 {
	divs := make([]float32, 0)

	Outer:
		for d := min; d <= max; d += 2 {
			
			for _, n := range nums {
				if math.Mod(float64(n), float64(d)) != 0 {
					continue Outer
				}
			}

			divs = append(divs, d)
		}

	return divs
}

/* Returns index of the first occurence of n within s or -1 if s does not contain n. */
func Index(s []float32, n float32) int {
	for i := range s {
		if s[i] == n {
			return i
		}
	}

	return -1
}
