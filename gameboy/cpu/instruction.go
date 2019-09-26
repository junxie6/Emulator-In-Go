package cpu

import "fmt"

var map_LD_Reg_n8bits = map[byte]registerID{
	0x06: B,
	0x0E: C,
	0x16: D,
	0x1E: E,
	0x26: H,
	0x2E: L,
}

// LD B,n  06 8
func (gb *GBCpu) instruction(opcode byte) (cycle int) {
	switch opcode {
	case 0x06, 0x0E, 0x16, 0x1E, 0x26, 0x2E:
		return gb._LD_Reg_n_8bits(map_LD_Reg_n8bits[opcode])

	}
	panic(fmt.Errorf("unknown opcode %x", opcode))
}

func (gb *GBCpu) _LD_Reg_n_8bits(x registerID) (cycle int) {

	n := gb.load8bits()
	rg := gb.registers[x]
	switch x {
	case PC, SP:
		rg.Write(uint16(n))
	default:
		if x >= Split {
			panic(fmt.Errorf("unknown register %d", x))
		} else {
			switch x % HiLoMagic {
			case 1:
				rg.WriteLo(n)
			case 0:
				rg.WriteHi(n)
			default:
				rg.Write(uint16(n))
			}
		}
	}
	return 8
}
