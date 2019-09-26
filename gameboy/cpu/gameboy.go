package cpu

import (
	"github.com/KeyOneLi/Emulator-In-Go/gameboy/util"
)

// GBCpu gameboy cpu instance
type GBCpu struct {
	memory    []byte
	registers map[registerID]*Register
	flag      byte
}

func NewGBCpu() *GBCpu {
	return &GBCpu{
		memory:    make([]byte, MemorySize),
		registers: preloadRegister(),
	}
}

func (gb *GBCpu) load8bits() byte {
	defer gb.increasePC()
	pc := gb.registers[PC]
	return gb.memory[pc.Read()]
}

func (gb *GBCpu) load16bits() uint16 {
	hi := gb.load8bits()
	lo := gb.load8bits()

	return util.ByteCombine(hi, lo)
}

func (gb *GBCpu) increasePC() {
	pc := gb.registers[PC]
	pc.Write(pc.Read() + 1)
}
