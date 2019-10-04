package cpu

import "io"

type CPU interface {
	LoadROM(rom io.Reader, offset int)
	LoadOpcode() byte
	ExcInstrution()
}
