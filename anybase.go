// Package anybase provides utilities for converting bytes to base-n
package anybase

type Anybase []byte

// Encode converts the source bytes to a destination base-n byte slice.
func (base Anybase) Encode(src, dst []byte) int {
	return -1
}

// Decode converts the destination bytes back into the byte slice.
func (base Anybase) Decode(dst, src []byte) (int, error) {
	return -1, nil
}

// EncodedLen returns the number of bytes needed to encode the source bytes.
func (base Anybase) EncodedLen(srcLen int) int {
	return -1
}

// DecodedLen returns the number of bytes that the decoded destination would
// be.
func (base Anybase) DecodedLen(dstLen int) int {
	return -1
}
