package crypt

func c2l(c []uint8) uint32 {
	return uint32(c[0]) | uint32(c[1])<<8 | uint32(c[2])<<16 | uint32(c[3])<<24
}

func l2c(l uint32, c []uint8) {
	c[0] = uint8(l & 0xff)
	c[1] = uint8((l >> 8) & 0xff)
	c[2] = uint8((l >> 16) & 0xff)
	c[3] = uint8((l >> 24) & 0xff)
	//log.Infof("l2c: l: %v c: (%x, %x, %x, %x)", l, c[0], c[1], c[2], c[3])
}

func PermOp(a uint32, b uint32, n int, m uint32) (retA uint32, retB uint32) {

	t := ((a >> n) ^ b) & m
	b ^= t
	a ^= t << n

	return a, b
}

func HPermOp(a uint32, n int, m uint32) (retA uint32) {
	t := ((a << (16 - n)) ^ a) & m
	a = a ^ t ^ (t >> (16 - n))
	return a
}
