package world

import (
	"crypto/sha256"
	"math/big"
	"math/rand"
	"strconv"
)

/* Checks whether the slice s contains the specified number n. */
func Contains(s []uint8, n uint8) bool {
	for _, e := range s {
		if e == n {
			return true
		}
	}
	return false
}

/* Returns the maximum value of int64. */
func maxInt64() int64 {
	return int64(^uint64(0) >> 1)
}

/* Returns new State buffer of specified size. */
func newStateBuffer(size int) []State {
	return make([]State, size, size)
}

/* Returns sha256sum of s. */
func sha256sum(s string) []byte {
	hash := sha256.New()
	hash.Write([]byte(s))
	return hash.Sum(nil)
}

/* 
	Seeds the PRNG with specified string. Numerical seeds are treated like integers. 
	For seeds containing other characters, seedFromString is called.
*/
func Seed(seed string) {
	n, err := strconv.ParseInt(seed, 10, 64)
	if err != nil {
		n = seedFromString(seed)
	}

	rand.Seed(n)
}

/* 
	Converts seed into int64. It is accomplished in the following way:
		1. sha256sum of the seed is calculated and saved as big.Int
		2. big.Int is reduced into int64 range with modulo maxInt64()
*/
func seedFromString(seed string) int64 {
	n := new(big.Int)
	n.SetBytes(sha256sum(seed))

	n.Mod(n, big.NewInt(maxInt64()))
	
	return n.Int64()
}
