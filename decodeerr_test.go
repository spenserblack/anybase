package anybase

import "testing"

func TestBadEncodedByte(t *testing.T) {
	err := ErrBadEncodedByte(128)
	want := "Invalid encoded byte: 128"

	if s := err.Error(); s != want {
		t.Fatalf(`err.Error() = %q, want %q`, s, want)
	}
}
