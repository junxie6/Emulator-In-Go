package cpu

type registerID int

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
	SP
	PC

	N  //8bits load
	NN //16bits load
)

type flag byte

const (
	flagZ byte = 0x80 >> iota
	flagN
	flagH
	flagC
)

const (
	MemorySize = 0x10000
)
