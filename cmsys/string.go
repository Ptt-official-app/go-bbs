package cmsys

func StringHash(theBytes []byte) uint32 {
	return fnv1a32StrCase(theBytes, FNV1_32_INIT)
}

func toupper(theByte byte) byte {
	if theByte >= 'a' && theByte <= 'z' {
		return theByte - 32
	}
	return theByte
}
