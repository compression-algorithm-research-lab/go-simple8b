package simple8b

import (
	"testing"
)

func TestDecodeE(t *testing.T) {

	slice := []int{
		1, 2, 3, 4, 100, -1, -2, -3, -4,
	}
	bytes := Encode(slice)
	t.Log(bytes)

	decodeSlice, err := DecodeE[int](bytes)
	t.Log(err)
	t.Log(decodeSlice)

}
