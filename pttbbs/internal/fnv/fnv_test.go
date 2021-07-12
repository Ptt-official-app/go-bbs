package fnv

import (
	"testing"
)

func TestFnv32a(t *testing.T) {
	type Case struct {
		input    string
		offset   uint32
		expected uint32
	}

	testcases := []Case{
		{
			input:    "",
			offset:   0x12345678,
			expected: 0x12345678,
		},
		{
			input:    "",
			offset:   0x811c9dc5,
			expected: 0x811c9dc5,
		},
		{
			// from golang golden32
			input:    "abc",
			offset:   0x811c9dc5,
			expected: 0x1a47e90b,
		},
		{
			input:    "12312",
			offset:   PttFnv32Init,
			expected: 0x7AEB94B2,
		},
		{
			input:    "PICHU",
			offset:   PttFnv32Init,
			expected: 0xA3389082,
		},
		{
			input:    "12312",
			offset:   12345,
			expected: 0x2d7500e0,
		},
	}

	for _, c := range testcases {
		h := New32aWith(c.offset)
		h.Write([]byte(c.input))
		b := []byte{}
		b = h.Sum(b)
		var actual uint32
		t.Log("b:", b)
		actual = uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
		if actual != c.expected {
			t.Errorf("fav32 hash failed, expected: 0x%X, got 0x%X", c.expected, actual)
		}

	}
}
