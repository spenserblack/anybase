package anybase

import "testing"

// TestBase2Decode tests that a base-2 implementation could be decoded.
func TestBase2Decode(t *testing.T) {
	base := Decoder{'x': 0, 'y': 1}
	src := []byte("xxxxxxxxxyxyxyxyyxyxyxyxyyyyyyyyyxxxxxxy")
	dst := make([]byte, 5)
	n, err := base.Decode(src, dst)
	if err != nil {
		t.Fatalf(`n = %v, want nil`, err)
	}
	if n != 5 {
		t.Fatalf(`n = %v, want 5`, n)
	}
	want := []byte{0, 0x55, 0xAA, 0xFF, 0b10000001}
	for i, want := range want {
		if dst[i] != want {
			t.Errorf(`dst[%d] = %v, want %v`, i, dst[i], want)
		}
	}
}

// TestBase3Decode tests that a base-3 implementation could be decoded.
func TestBase3Decode(t *testing.T) {
	decoder := Decoder{'0': 0, '1': 1, '2': 2}

	tests := []struct {
		dst  []byte
		want []byte
	}{
		{[]byte("100110"), []byte{0xFF}},
		{[]byte("000010"), []byte{3}},
		{[]byte("100110000010"), []byte{0xFF, 3}},
	}

	for _, tt := range tests {
		src := make([]byte, decoder.DecodedLen(len(tt.dst)))
		n, err := decoder.Decode(tt.dst, src)
		if err != nil {
			t.Fatalf(`err = %v, want nil`, err)
		}
		if n < len(tt.want) {
			t.Fatalf(`n = %v, want %v`, n, len(tt.want))
		}
		if string(src) != string(tt.want) {
			t.Errorf(`%s decoded with base-3 = %v, want %v`, tt.dst, src, tt.want)
		}
	}
}

// TestBase256Decode tests that a base-256 implementation could be decoded.
//

// NOTE Basically mapping a destination byte back to the source byte.
func TestBase256Decode(t *testing.T) {
	decoder := make(Decoder, 256)

	for i := byte(0); ; i++ {
		decoder[i] = 0xFF - i
		if i == 0xFF {
			break
		}
	}

	tests := []struct {
		dst  []byte
		want []byte
	}{
		{[]byte{0}, []byte{0xFF}},
		{[]byte{0xFF}, []byte{0}},
		{[]byte{0, 0xFF}, []byte{0xFF, 0}},
	}

	for _, tt := range tests {
		src := make([]byte, decoder.DecodedLen(len(tt.dst)))
		n, err := decoder.Decode(tt.dst, src)
		if err != nil {
			t.Fatalf(`err = %v, want nil`, err)
		}
		if n < len(tt.want) {
			t.Fatalf(`n = %v, want %v`, n, len(tt.want))
		}
		if string(src) != string(tt.want) {
			t.Errorf(`%s decoded with base-256 = %v, want %v`, tt.dst, src, tt.want)
		}
	}
}

func TestDecodedLen(t *testing.T) {
	tests := []struct {
		base   int
		dstLen int
		want   int
	}{
		{2, 8, 1},
		{2, 16, 2},
		{8, 3, 1},
		{8, 6, 2},
		{64, 2, 1},
		{128, 2, 1},
		{256, 1, 1},
		{3, 6, 1},
		{4, 4, 1},
	}

	for _, tt := range tests {
		decoder := make(Decoder, tt.base)
		for i := 0; i < tt.base; i++ {
			decoder[byte(i)] = byte(i)
		}

		if decodedLen := decoder.DecodedLen(tt.dstLen); decodedLen != tt.want {
			t.Errorf(
				`base-%d decoded length of %d bytes = %d, want %d`,
				tt.base,
				tt.dstLen,
				decodedLen,
				tt.want,
			)
		}
	}
}

// TestMakeEncoder tests that an encoder can be made from a decoder.
func TestMakeEncoder(t *testing.T) {
	decoder := Decoder{'a': 0, 'b': 1, 'c': 2}
	actual := decoder.Encoder()
	want := Encoder("abc")

	if actual, want := len(actual), len(want); actual != want {
		t.Fatalf(`len(actual) = %d, want %d`, actual, want)
	}

	for i, want := range want {
		if actual := actual[i]; actual != want {
			t.Errorf(`actual[%d] = %v, want %v`, i, actual, want)
		}
	}
}

// TestDecoderBadDst tests that the decoder returns an error when the dst
// contains invalid bytes.
func TestDecoderBadDst(t *testing.T) {
	decoder := make(Decoder, 16)
	chars := "0123456789abcdef"
	for i := 0; i < len(chars); i++ {
		decoder[chars[i]] = byte(i)
	}
	want := ErrBadEncodedByte('g')

	if _, err := decoder.Decode([]byte{'f', 'g'}, make([]byte, 1)); err != want {
		t.Fatalf(`err = %v, want %v`, err, want)
	}
}
