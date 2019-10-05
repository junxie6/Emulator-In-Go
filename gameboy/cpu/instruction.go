package cpu

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

func (gb *GBCpu) _LD(opcode byte, params []registerID) {
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
		if n < 0 {
			// TODO: carry flag?
			value = gb.registers.Get(PC).Read() - uint16(-n)
		} else {
			value = uint16(n) + gb.registers.Get(PC).Read()
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

func (gb *GBCpu) _8bitsLDr1r2(param paramWraper) (cycle int) {
	panic("to do")
	//return param.cycles
}

func (gb *GBCpu) _8bitsLDAorC(param paramWraper) (cycle int) {
	panic("to do")
}
