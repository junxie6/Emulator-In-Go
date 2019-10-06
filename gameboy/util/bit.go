package util

//ByteCombine trans 2 8bits number to 16bit number
func ByteCombine(hi, lo byte) uint16 {
	return (uint16(hi) << 8) | uint16(lo)
}

//HalfCarryForAdd check if half carry happened for a+b
func HalfCarryForAdd(a, b uint16) bool {
	return (a&0xf)+(b&0xf) > 0xf
}

//HalfCarryForSub check if half carry happened for a-b
func HalfCarryForSub(a, b uint16) bool {
	return (a & 0xf) < (b & 0xf)
}

//CarryForAdd check if carry happened for a+b
func CarryForAdd(a, b uint16) bool {
	return (uint16(a)&0xf)+(uint16(b)&0xf) > 0xf
}

//CarryForSub check if carry happend for a-b
func CarryForSub(a, b uint16) bool {
	return (a & 0xff) < (b & 0xff)
}
