package main

import "math"

/* Returns the member of s closest to the mean of all members. Returns 0 given an empty slice. */
func GetClosestToMean(s []float32) float32 {
	if len(s) == 0 {
		return 0
	}

	mean := Sum(s) / float32(len(s))

	closest := s[0]

	for i := range s[1:] {
		if math.Abs(float64(s[i]-mean)) < math.Abs(float64(closest-mean)) {
			closest = s[i]
		}
	}

	return closest
}

/* Returns a slice containing common divisors of numbers in nums between min and max. */
func GetCommonDivisors(min, max float32, nums ...float32) []float32 {
	divs := make([]float32, 0)

Outer:
	for d := min; d <= max; d++ {

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

/* Returns sum of the numbers in s. */
func Sum(s []float32) (sum float32) {
	for i := range s {
		sum += s[i]
	}

	return
}
