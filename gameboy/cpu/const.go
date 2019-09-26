package cpu

type registerID int

const HiLoMagic = 3

const (
	A registerID = iota
	F
	AF
	B
	C
	BC
	D
	E
	DE
	H
	L
	HL
	Split
	SP
	PC
)

const (
	MemorySize = 0x10000
)
