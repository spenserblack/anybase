package anybase

import "testing"

// TestBase2Encode tests that a base-2 implementation could be encoded.
func TestBase2Encode(t *testing.T) {
	base := Encoder("xy")
	src := []byte{0, 0x55, 0xAA, 0xFF, 0b10000001}
	dst := make([]byte, 40)
	n := base.Encode(src, dst)
	if n != 40 {
		t.Fatalf(`n = %v, want 40`, n)
	}
	want := []byte("xxxxxxxxxyxyxyxyyxyxyxyxyyyyyyyyyxxxxxxy")
	for i, want := range want {
		if dst[i] != want {
			t.Errorf(`dst[%d] = %c, want %c`, i, dst[i], want)
		}
	}
}

// TestBase3Encode tests that a base-3 implementation could be encoded.
func TestBase3Encode(t *testing.T) {
	base := Encoder("012")

	tests := []struct {
		src  []byte
		want []byte
	}{
		{[]byte{0xFF}, []byte("100110")},
		{[]byte{3}, []byte("000010")},
		{[]byte{0xFF, 3}, []byte("100110000010")},
	}

	for _, tt := range tests {
		dst := make([]byte, base.EncodedLen(len(tt.src)))
		if n := base.Encode(tt.src, dst); n < len(tt.want) {
			t.Fatalf(`n = %v, want %v`, n, len(tt.want))
		}
		if string(dst) != string(tt.want) {
			t.Errorf(`%v encoded with base-3 = %s, want %s`, tt.src, dst, tt.want)
		}
	}
}

// TestBase256Encode tests that a base-256 implementation could be encoded.
//
// NOTE Base-256 basically just maps one byte to another.
func TestBase256Encode(t *testing.T) {
	encoder := make(Encoder, 0, 256)

	for i := byte(0); ; i++ {
		encoder = append(encoder, 0xFF-i)
		if i == 0xFF {
			break
		}
	}

	tests := []struct {
		src  []byte
		want []byte
	}{
		{[]byte{0xFF}, []byte{0}},
		{[]byte{0}, []byte{0xFF}},
		{[]byte{0xFF, 0}, []byte{0, 0xFF}},
	}

	for _, tt := range tests {
		dst := make([]byte, encoder.EncodedLen(len(tt.src)))
		if n := encoder.Encode(tt.src, dst); n < len(tt.want) {
			t.Fatalf(`n = %v, want %v`, n, len(tt.want))
		}
		if string(dst) != string(tt.want) {
			t.Errorf(`%v encoded with base-256 = %s, want %s`, tt.src, dst, tt.want)
		}
	}
}

func TestEncodedLen(t *testing.T) {
	tests := []struct {
		base   int
		srcLen int
		want   int
	}{
		{2, 1, 8},
		{2, 2, 16},
		{8, 1, 3},
		{8, 2, 6},
		{64, 1, 2},
		{128, 1, 2},
		{256, 1, 1},
		{3, 1, 6},
		{4, 1, 4},
	}

	for _, tt := range tests {
		encoder := make(Encoder, tt.base)

		if encodedLen := encoder.EncodedLen(tt.srcLen); encodedLen != tt.want {
			t.Errorf(
				`base-%d encoded length of %d bytes = %d, want %d`,
				tt.base,
				tt.srcLen,
				encodedLen,
				tt.want,
			)
		}
	}
}

// TestMakeDecoder tests that a decoder can be made from an encoder.
func TestMakeDecoder(t *testing.T) {
	s := "0123456789abcdef"
	encoder := Encoder(s)
	decoder := encoder.Decoder()

	if l := len(decoder); l != 16 {
		t.Fatalf(`len(decoder) = %v, want 16`, l)
	}

	for i := 0; i < len(s); i++ {
		b := s[i]
		want := byte(i)
		if actual, ok := decoder[b]; !ok || actual != want {
			t.Errorf(`[exists: %v] decoder[%v] = %v, want %v`, ok, b, actual, want)
		}
	}
}
