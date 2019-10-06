package util

//ByteCombine trans 2 8bits number to 16bit number
func ByteCombine(hi, lo byte) uint16 {
	return (uint16(hi) << 8) | uint16(lo)
}

//HalfCarry check if half carry happened for a+b
func HalfCarry(a, b uint16) bool {
	return (a&0xf)+(b&0xf) > 0xf
}

//HalfBorrow check if half carry happened for a-b
func HalfBorrow(a, b uint16) bool {
	return (a & 0xf) < (b & 0xf)
}

//Carry check if carry happened for a+b
func Carry(a, b uint16) bool {
	return (uint16(a)&0xf)+(uint16(b)&0xf) > 0xf
}

//Borrow check if carry happend for a-b
func Borrow(a, b uint16) bool {
	return (a & 0xff) < (b & 0xff)
}
