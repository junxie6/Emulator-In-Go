package cpu

import (
	"github.com/KeyOneLi/Emulator-In-Go/gameboy/util"
)

// GBCpu gameboy cpu instance
type GBCpu struct {
	memory    Memory
	registers *RegisterPool
	flag      byte
}

func NewGBCpu() *GBCpu {
	return &GBCpu{
		// memory:    make([]byte, MemorySize),
		registers: NewRegisterPool(),
	}
}

func (gb *GBCpu) load8bits() byte {
	defer gb.increasePC()
	pc := gb.registers.Get(PC)
	return gb.memory.ReadAt(pc.Read())
}

func (gb *GBCpu) load16bits() uint16 {
	hi := gb.load8bits()
	lo := gb.load8bits()

	return util.ByteCombine(hi, lo)
}

func (gb *GBCpu) increasePC() {
	pc := gb.registers.Get(PC)
	pc.Write(pc.Read() + 1)
}
