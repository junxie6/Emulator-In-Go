package util

//ByteCombine trans 2 8bits number to 16bit number
func ByteCombine(hi, lo byte) uint16 {
	return (uint16(hi) << 8) | uint16(lo)
}
