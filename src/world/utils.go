package world

/* Checks whether the slice s contains the specified number n. */
func Contains(s []uint8, n uint8) bool {
	for _, e := range s {
		if e == n {
			return true
		}
	}
	return false
}
