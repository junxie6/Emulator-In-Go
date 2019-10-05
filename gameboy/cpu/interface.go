package cpu

import "io"

type CPU interface {
	LoadROM(rom io.Reader, offset int)
	LoadOpcode() byte
	ExcInstrution()
}

type Memory interface {
	ReadAt(addr uint16) byte
	WriteAt(addr uint16, value byte)
}
