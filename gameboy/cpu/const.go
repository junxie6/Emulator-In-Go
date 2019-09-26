package cpu

type registerID int

const (
	AF registerID = iota
	BC
	DE
	HL
	SP
	PC
)

const (
	MemorySize = 0x10000
)
