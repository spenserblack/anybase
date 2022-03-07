// Package anybase provides utilities for converting bytes to base-n
package anybase

// DigitsForByte returns the number if digits used for a single byte.
func digitsForByte(digitCount int) int {
	// TODO Optimize
	i := 0
	v := 0xFF
	for ; v > 0; i++ {
		v /= digitCount
	}
	return i
}
