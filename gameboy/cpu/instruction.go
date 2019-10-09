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
			flagHalfCarry, flagCarry = util.HalfBorrow(a, b), util.Borrow(a, b)
		} else {
			a := gb.registers.Get(SP).Read()
			b := uint16(n)
			value = a + b
			flagHalfCarry, flagCarry = util.HalfCarry(a, b), util.Carry(a, b)
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

	if util.HalfCarry(a, b) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}

	if util.Carry(a, b) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}
}

func (gb *GBCpu) adc(opcode byte, params []registerID) {
	var carry, n uint16
	p2 := params[1]

	if gb.registers.GetFlag(flagC) {
		carry = 1
	}

	switch p2 {
	case N:
		n = uint16(gb.load8bits())
	case NN:
		panic("invalid")
	default:
		n = gb.registers.Get(p2).Read()
	}

	a := gb.registers.Get(A).Read()
	gb.registers.Get(A).Write(a + n + carry)

	gb.registers.ResetFlag(flagN)
	if a+n+carry == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	if util.HalfCarry(a, n) || util.HalfCarry(a+n, carry) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}

	if util.Carry(a, n) || util.Carry(a+n, carry) {
		gb.registers.SetFlag(flagC)
	} else {
		gb.registers.ResetFlag(flagC)
	}
}

func (gb *GBCpu) sub(opcode byte, params []registerID) {
	var a, b uint16

	a = gb.registers.Get(A).Read()
	switch params[1] {
	case N:
		b = uint16(gb.load8bits())
	case NN:
		//invalid?
		panic("invalid")
	default:
		b = gb.registers.Get(params[1]).Read()
	}

	value := a - b
	gb.registers.Get(A).Write(value)

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	gb.registers.SetFlag(flagN)

	if util.HalfBorrow(a, b) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}

	if util.Borrow(a, b) {
		gb.registers.SetFlag(flagC)
	} else {
		gb.registers.ResetFlag(flagC)
	}
}

func (gb *GBCpu) sbc(opcode byte, params []registerID) {

	var carry uint16
	if gb.registers.GetFlag(flagC) {
		carry = 1
	}

	var n uint16
	a := gb.registers.Get(A).Read()

	switch params[1] {
	case N:
		n = uint16(gb.load8bits())
	case NN:
		panic("invalid")
	default:
		n = gb.registers.Get(params[1]).Read()
	}

	value := a - n - carry
	gb.registers.Get(A).Write(value)

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	gb.registers.SetFlag(flagN)

	if util.HalfBorrow(a, n) || util.HalfBorrow(a-n, carry) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}

	if util.Borrow(a, n) || util.Borrow(a-n, carry) {
		gb.registers.SetFlag(flagC)
	} else {
		gb.registers.ResetFlag(flagC)
	}
}

func (gb *GBCpu) and(opcode byte, params []registerID) {
	var a, b uint16
	a = gb.registers.Get(A).Read()
	switch params[0] {
	case N:
		b = uint16(gb.load8bits())
	case NN:
		panic("invalid")
	default:
		b = gb.registers.Get(params[1]).Read()
	}

	value := a & b
	gb.registers.Get(A).Write(value)

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagC)
	gb.registers.SetFlag(flagH)
}

func (gb *GBCpu) or(opcode byte, params []registerID) {
	var a, b uint16
	a = gb.registers.Get(A).Read()
	switch params[0] {
	case N:
		b = uint16(gb.load8bits())
	case NN:
		panic("invalid")
	default:
		b = gb.registers.Get(params[1]).Read()
	}

	value := a | b
	gb.registers.Get(A).Write(value)

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagC)
	gb.registers.ResetFlag(flagH)
}

func (gb *GBCpu) xor(opcode byte, params []registerID) {
	var a, b uint16
	a = gb.registers.Get(A).Read()
	switch params[0] {
	case N:
		b = uint16(gb.load8bits())
	case NN:
		panic("invalid")
	default:
		b = gb.registers.Get(params[1]).Read()
	}

	value := a ^ b
	gb.registers.Get(A).Write(value)

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagC)
	gb.registers.ResetFlag(flagH)
}

func (gb *GBCpu) cp(opcode byte, params []registerID) {
	var a, n uint16

	a = gb.registers.Get(A).Read()
	switch params[0] {
	case N:
		n = uint16(gb.load8bits())
	case NN:
		panic("invalid")
	default:
		n = gb.registers.Get(params[0]).Read()
	}

	value := a - n
	gb.registers.Get(A).Write(value)

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}

	gb.registers.SetFlag(flagN)

	if util.HalfBorrow(a, n) {
		gb.registers.ResetFlag(flagH)
	} else {
		gb.registers.SetFlag(flagH)
	}

	if util.Borrow(a, n) {
		gb.registers.ResetFlag(flagC)
	} else {
		gb.registers.SetFlag(flagC)
	}
}

func (gb *GBCpu) inc(opcode byte, params []registerID) {
	r := gb.registers.Get(params[0])
	n := r.Read()
	r.Write(n + 1)

	if n+1 == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}
	gb.registers.ResetFlag(flagN)

	if util.HalfBorrow(n, 1) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}
}

func (gb *GBCpu) dec(opcode byte, params []registerID) {
	r := gb.registers.Get(params[0])
	n := r.Read()
	r.Write(n - 1)

	if n-1 == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}
	gb.registers.SetFlag(flagN)

	if !util.HalfBorrow(n, 1) {
		gb.registers.SetFlag(flagH)
	} else {
		gb.registers.ResetFlag(flagH)
	}
}

func (gb *GBCpu) swap(opcode byte, params []registerID) {
	panic("to do")
}

func (gb *GBCpu) daa(opcode byte, params []registerID) {

}

func (gb *GBCpu) cpl(opcode byte, params []registerID) {
	a := gb.registers.Get(A)

	a.Write(^a.Read())

	gb.registers.SetFlag(flagN)
	gb.registers.SetFlag(flagH)
}

func (gb *GBCpu) ccf(opcode byte, params []registerID) {
	if gb.registers.GetFlag(flagC) {
		gb.registers.ResetFlag(flagC)
	} else {
		gb.registers.SetFlag(flagC)
	}

	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagH)
}

func (gb *GBCpu) scf(opcode byte, params []registerID) {
	gb.registers.SetFlag(flagC)
	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagH)
}

func (gb *GBCpu) nop(opcode byte, params []registerID) {
	//nop
}

func (gb *GBCpu) halt(opcode byte, params []registerID) {
	panic("to do???")
}

func (gb *GBCpu) stop(opcode byte, params []registerID) {
	panic("to do???")
}

// disable interrupts
func (gb *GBCpu) di(opcode byte, params []registerID) {
	panic("to do")
}

// enable interrupts
func (gb *GBCpu) ei(opcode byte, params []registerID) {
	panic("to do")
}

func (gb *GBCpu) rlca(opcode byte, params []registerID) {
	a := gb.registers.Get(A)

	value := byte(a.Read())
	oldBit7 := (value & 0x80) >> 7

	value = (value << 1) | oldBit7

	a.Write(uint16(value))

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}
	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagH)

	if oldBit7 == 0 {
		gb.registers.ResetFlag(flagC)
	} else {
		gb.registers.SetFlag(flagC)
	}
}

func (gb *GBCpu) rla(opcode byte, params []registerID) {

	var newBit0 byte = 0

	if gb.registers.GetFlag(flagC) {
		newBit0 = 1
	}
	a := gb.registers.Get(A)

	value := byte(a.Read())
	oldBit7 := (value & 0x80) >> 7

	value = (value << 1) | newBit0

	a.Write(uint16(value))

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}
	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagH)

	if oldBit7 == 0 {
		gb.registers.ResetFlag(flagC)
	} else {
		gb.registers.SetFlag(flagC)
	}
}

func (gb *GBCpu) rrca(opcode byte, params []registerID) {
	a := gb.registers.Get(A)

	value := byte(a.Read())
	oldBit0 := (value & 0x1)

	value = (value >> 1) | (oldBit0 << 7)

	a.Write(uint16(value))

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}
	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagH)

	if oldBit0 == 0 {
		gb.registers.ResetFlag(flagC)
	} else {
		gb.registers.SetFlag(flagC)
	}
}

func (gb *GBCpu) rra(opcode byte, params []registerID) {

	var newBit7 byte = 0

	if gb.registers.GetFlag(flagC) {
		newBit7 = 1
	}
	a := gb.registers.Get(A)

	value := byte(a.Read())
	oldBit0 := (value & 0x1)

	value = (value >> 1) | (newBit7 << 7)

	a.Write(uint16(value))

	if value == 0 {
		gb.registers.SetFlag(flagZ)
	} else {
		gb.registers.ResetFlag(flagZ)
	}
	gb.registers.ResetFlag(flagN)
	gb.registers.ResetFlag(flagH)

	if oldBit0 == 0 {
		gb.registers.ResetFlag(flagC)
	} else {
		gb.registers.SetFlag(flagC)
	}
}

func (gb *GBCpu) rlc(opcode byte, params []registerID) {
	panic("to do")
}
