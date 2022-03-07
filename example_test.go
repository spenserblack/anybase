package anybase_test

import (
	"fmt"

	"github.com/spenserblack/anybase"
)

func Example() {
	encoder := anybase.Encoder("0123456789abcdef")
	// It will often be useful to create an encoder-decoder pair
	decoder := encoder.Decoder()
	src1 := []byte{0xAB, 0xCD}
	dst := make([]byte, encoder.EncodedLen(len(src1)))
	encoder.Encode(src1, dst)
	fmt.Printf("%s\n", dst)

	src2 := make([]byte, decoder.DecodedLen(len(dst)))
	decoder.Decode(dst, src2)
	fmt.Printf("[%X, %X]\n", src2[0], src2[1])
	// Output:
	// abcd
	// [AB, CD]
}

func ExampleEncoder() {
	// Base-2, where the digits are "a" and "b"
	encoder := anybase.Encoder("ab")
	src := []byte{0b01011001}
	dst := make([]byte, encoder.EncodedLen(len(src)))
	encoder.Encode(src, dst)
	fmt.Printf("%s", dst)
	// Output: ababbaab
}

func ExampleDecoder() {
	// Base-3, where the digits are 1, 2, and 3.
	decoder := anybase.Decoder{
		'1': 0,
		'2': 1,
		'3': 2,
	}
	dst := []byte("111123")
	src := make([]byte, decoder.DecodedLen(len(dst)))
	decoder.Decode(dst, src)
	fmt.Printf("%v", src)
	// Output: [5]
}
