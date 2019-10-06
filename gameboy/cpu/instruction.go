package cpu

import (
	"github.com/KeyOneLi/Emulator-In-Go/gameboy/util"
)

// InstructionInfo shows detail info of an instruction
type InstructionInfo struct {
	Opcode      string   `json:"opcode"`
	Instruction string   `json:"instruction"`
	Params      []string `json:"params"`
	Code        byte     `json:"code"`
	Cycles      int      `json:"cycles"`
}

func (in *InstructionInfo) GetParams() []registerID {
	rtn := []registerID{}

	for _, raw := range in.Params {
		rtn = append(rtn, atoi(raw))
	}
	return rtn
}

// LD B,n  06 8
func (gb *GBCpu) instructions(opcode byte) (cycle int) {
	panic("to do")
}

func (gb *GBCpu) ld(opcode byte, params []registerID) {
	p1, p2 := params[0], params[1]
	var value uint16
	switch opcode {
	// LD A,(C)
	case 0xf2:
		value = uint16(gb.memory.ReadAt(0xff00)) + gb.registers.Get(C).Read()
		gb.registers.Get(A).Write(value)
	// LD (C),A
	case 0xe2:
		value = gb.registers.Get(A).Read()
		addr := 0xff00 + gb.registers.Get(C).Read()
		gb.memory.WriteAt(addr, byte(value))
	// LDH (n), A
	case 0xe0:
		value = gb.registers.Get(A).Read()
		addr := 0xff00 + uint16(gb.load8bits())
		gb.memory.WriteAt(addr, byte(value))
	// LDH A,(n)
	case 0xf0:
		value = 0xff00 + uint16(gb.load8bits())
		gb.registers.Get(A).Write(value)
	case 0xf8:
		n := int8(gb.load8bits())
		gb.registers.ResetFlag(flagZ)
		gb.registers.ResetFlag(flagN)

		flagHalfCarry, flagCarry := false, false

		if n < 0 {
			a := gb.registers.Get(SP).Read()
			b := uint16(-n)
			value = a - b
			flagHalfCarry, flagCarry = util.HalfCarryForSub(a, b), util.CarryForSub(a, b)
		} else {
			a := gb.registers.Get(SP).Read()
			b := uint16(n)
			value = a + b
			flagHalfCarry, flagCarry = util.HalfCarryForAdd(a, b), util.CarryForAdd(a, b)
		}

		if flagCarry {
			gb.registers.SetFlag(flagC)
		} else {
			gb.registers.ResetFlag(flagC)
		}

		if flagHalfCarry {
			gb.registers.SetFlag(flagH)
		} else {
			gb.registers.ResetFlag(flagH)
		}

		gb.registers.Get(HL).Write(value)
	case 0x08:
		nn := gb.load16bits()
		gb.registers.Get(SP).Write(nn)
	default:
		switch p2 {
		case N:
			value = uint16(gb.load8bits())
		case NN:
			value = gb.load16bits()
		default:
			value = gb.registers.Get(p2).Read()
		}

		switch p1 {
		case NN:
			nn := gb.load16bits()
			gb.memory.WriteAt(nn, byte(value))
		case N:
			panic("I think it is invaild case")
		default:
			gb.registers.Get(p1).Write(value)
		}
	}

	switch opcode {

	// LDD A,(HL) or LDD (HL), A: post-decreatement
	case 0x3a, 0x32:
		value = gb.registers.Get(HL).Read() - 1
		gb.registers.Get(HL).Write(value)
	// LDI A,(HL) or LDI (HL), A: post-increatement
	case 0x2a, 0x22:
		value = gb.registers.Get(HL).Read() + 1
		gb.registers.Get(HL).Write(value)
	}
}

func (gb *GBCpu) push(opcode byte, params []registerID) {
	sp := gb.registers.Get(SP).Read()
	r := gb.registers.Get(params[0]).(*Register16bit)

	gb.memory.WriteAt(sp, r.ReadLo())
	gb.memory.WriteAt(sp-1, r.ReadHi())
	gb.registers.Get(SP).Write(sp - 2)
}

func (gb *GBCpu) pop(opcode byte, params []registerID) {
	sp := gb.registers.Get(SP).Read()

	hi := gb.memory.ReadAt(sp)
	lo := gb.memory.ReadAt(sp + 1)

	value := util.ByteCombine(hi, lo)

	gb.registers.Get(params[0]).Write(value)
	gb.registers.Get(SP).Write(sp + 2)
}

func (gb *GBCpu) add(opcode byte, params []registerID) {
	// p1 := params[0]

	p2 := params[1]
	var a, b uint16

	a = gb.registers.Get(A).Read()
	switch p2 {
	case N:
		b = uint16(gb.load8bits())
	case NN:
		//invalid?
		panic("invalid")
	default:
		b = gb.registers.Get(p2).Read()
	}
	value := a + b
	gb.registers.Get(A).Write(value)

	gb.registers.ResetFlag(flagN)
	if value&0xff == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	if util.HalfCarryForAdd(a, b) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}

	if util.CarryForAdd(a, b) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}
}

func (gb *GBCpu) adc(opcode byte, params []registerID) {
	var carry uint16
	p2 := params[1]

	if gb.registers.GetFlag(flagC) {
		carry = 1
	}

	n := gb.registers.Get(p2).Read()
	a := gb.registers.Get(A).Read()
	gb.registers.Get(A).Write(a + n + carry)

	gb.registers.ResetFlag(flagN)
	if a+n+carry == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	if util.HalfCarryForAdd(a, n) || util.HalfCarryForAdd(a+n, carry) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}

	if util.CarryForAdd(a, n) || util.CarryForAdd(a+n, carry) {
		gb.registers.SetFlag(flagC)
	} else {
		gb.registers.ResetFlag(flagC)
	}
}
