package anybase

// Encoder encodes source bytes to a destination byte slice.
type Encoder []byte

// Encode converts the source bytes to a destination base-n byte slice.
func (encoder Encoder) Encode(src, dst []byte) int {
	length := len(encoder)
	if length < 2 {
		panic("An encoder must have at least 2 bytes")
	}
	if length > 256 {
		panic("An encoder cannot have more than 256 bytes")
	}
	digitsForByte := digitsForByte(length)
	count := 0
	for srcIndex, srcByte := range src {
		srcByte := int(srcByte)
		for i := 0; i < digitsForByte; i++ {
			// NOTE Index from right to put bytes in proper order.
			rightIndex := digitsForByte - i - 1
			dst[rightIndex+srcIndex*digitsForByte] = encoder[srcByte%length]
			srcByte /= length
			count++
		}
	}
	return count
}

// EncodedLen returns the length of an encoded dst.
func (encoder Encoder) EncodedLen(srcLen int) int {
	return digitsForByte(len(encoder)) * srcLen
}

// Decoder creates a decoder that is compatible with the encoder.
func (encoder Encoder) Decoder() Decoder {
	decoder := make(Decoder, len(encoder))

	for i, v := range encoder {
		decoder[v] = byte(i)
	}
	return decoder
}
