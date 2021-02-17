package fnv

import (
	"encoding"
	"hash"
	"hash/fnv"
)

const (
	PttFnv32Init uint32 = 33554467
)

func New32aWith(offset uint32) hash.Hash32 {

	h := fnv.New32a()
	b, _ := h.(encoding.BinaryMarshaler).MarshalBinary()
	b[4] = byte((offset >> 24) & 0xFF)
	b[5] = byte(offset >> 16 & 0xFF)
	b[6] = byte(offset >> 8 & 0xFF)
	b[7] = byte(offset & 0xFF)

	h.(encoding.BinaryUnmarshaler).UnmarshalBinary(b)
	return h
}
