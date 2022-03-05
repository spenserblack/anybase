package anybase

import "fmt"

// Decoder decodes from the source bytes to a destination byte slice.
type Decoder map[byte]byte

// ErrBadEncodedByte is an error type for an invalid encoded byte.
type ErrBadEncodedByte byte

// Decode converts the destination bytes back into the byte slice.
func (decoder Decoder) Decode(dst, src []byte) (int, error) {
	count := 0
	length := len(decoder)
	digitsForByte := digitsForByte(length)
	for dstIndex, dstByte := range dst {
		decodedVal, ok := decoder[dstByte]
		if !ok {
			return count, ErrBadEncodedByte(dstByte)
		}
		// NOTE The significance of the encoded char
		exponent := digitsForByte - (dstIndex % digitsForByte) - 1
		// NOTE No worries about overflow when converting length to byte,
		// because exponent will be 0 with 256 length, resulting in 1.
		src[dstIndex/digitsForByte] += decodedVal * pow(byte(length), exponent)
		if dstIndex%digitsForByte == digitsForByte-1 {
			count++
		}
	}
	return count, nil
}

// DecodedLen returns the number of bytes that the decoded destination would
// be.
func (decoder Decoder) DecodedLen(dstLen int) int {
	return dstLen / digitsForByte(len(decoder))
}

// Encoder returns an encoder that is compatible with the decoder.
func (decoder Decoder) Encoder() Encoder {
	encoder := make(Encoder, len(decoder))

	for k, v := range decoder {
		encoder[v] = k
	}
	return encoder
}

func (err ErrBadEncodedByte) Error() string {
	return fmt.Sprintf("Invalid encoded byte: %v", byte(err))
}

func pow(n byte, exp int) byte {
	// NOTE We don't care about negative exponents
	if exp <= 0 {
		return 1
	}
	return n * pow(n, exp-1)
}
