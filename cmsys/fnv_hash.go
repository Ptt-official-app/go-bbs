package cmsys

//https://github.com/ptt/pttbbs/blob/master/include/fnv_hash.h
//commit: 6bdd36898bde207683a441cdffe2981e95de5b20

func fnv32Bytes(theBytes []byte, hval uint32) uint32 {
	for _, each := range theBytes {
		hval *= FNV_32_PRIME
		hval ^= uint32(each)
	}

	return hval
}

func fnv1a32Bytes(theBytes []byte, hval uint32) uint32 {
	for _, each := range theBytes {
		if each == 0 {
			break
		}
		hval ^= uint32(each)
		hval *= FNV_32_PRIME
	}

	return hval
}

func fnv1a32StrCase(theBytes []byte, hval uint32) uint32 {
	for _, each := range theBytes {
		if each == 0 {
			break
		}
		hval ^= uint32(toupper(each))
		hval *= FNV_32_PRIME
	}

	return hval
}

func fnv1a32DBCSCase(theBytes []byte, hval uint32) uint32 {
	isDBCS := false
	for _, each := range theBytes {
		if each == 0 {
			break
		}
		if isDBCS {
			// 2nd DBCS
			isDBCS = false
		} else {
			if each < 0x80 {
				each = toupper(each)
			} else {
				isDBCS = true
			}

		}
		hval ^= uint32(each)
		hval *= FNV_32_PRIME
	}

	return hval
}

//////////
//64bits
//////////

func fnv64Bytes(theBytes []byte, hval uint64) uint64 {
	for _, each := range theBytes {
		hval *= FNV_64_PRIME
		hval ^= uint64(each)
	}

	return hval
}

func fnv1a64Bytes(theBytes []byte, hval uint64) uint64 {
	for _, each := range theBytes {
		if each == 0 {
			break
		}
		hval ^= uint64(each)
		hval *= FNV_64_PRIME
	}

	return hval
}

func fnv1a64StrCase(theBytes []byte, hval uint64) uint64 {
	for _, each := range theBytes {
		if each == 0 {
			break
		}
		hval ^= uint64(toupper(each))
		hval *= FNV_64_PRIME
	}

	return hval
}

func fnv1a64DBCSCase(theBytes []byte, hval uint64) uint64 {
	isDBCS := false
	for _, each := range theBytes {
		if each == 0 {
			break
		}
		if isDBCS {
			// 2nd DBCS
			isDBCS = false
		} else {
			if each < 0x80 {
				each = toupper(each)
			} else {
				isDBCS = true
			}

		}
		hval ^= uint64(each)
		hval *= FNV_64_PRIME
	}

	return hval
}

func fnv1aByte(theByte byte, hval uint32) uint32 {
	hval ^= uint32(theByte)
	hval *= FNV_32_PRIME
	return hval
}
