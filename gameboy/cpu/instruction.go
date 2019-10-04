package cpu

import "fmt"

var mapReg_ld_nn_n8bits = map[byte]registerID{}

type paramWraper struct {
	params []registerID
	cycles int
}

var map8bitsLDr1r2 = map[byte]paramWraper{
	0x06: paramWraper{[]registerID{B, N}, 8},
	0x0E: paramWraper{[]registerID{C, N}, 8},
	0x16: paramWraper{[]registerID{D, N}, 8},
	0x1E: paramWraper{[]registerID{E, N}, 8},
	0x26: paramWraper{[]registerID{H, N}, 8},
	0x2E: paramWraper{[]registerID{L, N}, 8},

	0x7f: paramWraper{[]registerID{A, A}, 4},
	0x78: paramWraper{[]registerID{A, B}, 4},
	0x79: paramWraper{[]registerID{A, C}, 4},
	0x7A: paramWraper{[]registerID{A, D}, 4},
	0x7B: paramWraper{[]registerID{A, E}, 4},
	0x7C: paramWraper{[]registerID{A, H}, 4},
	0x7D: paramWraper{[]registerID{A, L}, 4},
	0x7E: paramWraper{[]registerID{A, HL}, 8},

	0x40: paramWraper{[]registerID{B, B}, 4},
	0x41: paramWraper{[]registerID{B, C}, 4},
	0x42: paramWraper{[]registerID{B, D}, 4},
	0x43: paramWraper{[]registerID{B, E}, 4},
	0x44: paramWraper{[]registerID{B, H}, 4},
	0x45: paramWraper{[]registerID{B, L}, 4},
	0x46: paramWraper{[]registerID{B, HL}, 8},
	0x47: paramWraper{[]registerID{B, A}, 4},

	0x4f: paramWraper{[]registerID{C, A}, 4},
	0x48: paramWraper{[]registerID{C, B}, 4},
	0x49: paramWraper{[]registerID{C, C}, 4},
	0x4a: paramWraper{[]registerID{C, D}, 4},
	0x4b: paramWraper{[]registerID{C, E}, 4},
	0x4c: paramWraper{[]registerID{C, H}, 4},
	0x4d: paramWraper{[]registerID{C, L}, 4},
	0x4e: paramWraper{[]registerID{C, HL}, 8},

	0x57: paramWraper{[]registerID{D, A}, 4},
	0x50: paramWraper{[]registerID{D, B}, 4},
	0x51: paramWraper{[]registerID{D, C}, 4},
	0x52: paramWraper{[]registerID{D, D}, 4},
	0x53: paramWraper{[]registerID{D, E}, 4},
	0x54: paramWraper{[]registerID{D, H}, 4},
	0x55: paramWraper{[]registerID{D, L}, 4},
	0x56: paramWraper{[]registerID{D, HL}, 8},

	0x5f: paramWraper{[]registerID{E, A}, 4},
	0x58: paramWraper{[]registerID{E, B}, 4},
	0x59: paramWraper{[]registerID{E, C}, 4},
	0x5a: paramWraper{[]registerID{E, D}, 4},
	0x5b: paramWraper{[]registerID{E, E}, 4},
	0x5c: paramWraper{[]registerID{E, H}, 4},
	0x5d: paramWraper{[]registerID{E, L}, 4},
	0x5e: paramWraper{[]registerID{E, HL}, 8},

	0x67: paramWraper{[]registerID{H, A}, 4},
	0x60: paramWraper{[]registerID{H, B}, 4},
	0x61: paramWraper{[]registerID{H, C}, 4},
	0x62: paramWraper{[]registerID{H, D}, 4},
	0x63: paramWraper{[]registerID{H, E}, 4},
	0x64: paramWraper{[]registerID{H, H}, 4},
	0x65: paramWraper{[]registerID{H, L}, 4},
	0x66: paramWraper{[]registerID{H, HL}, 8},

	0x6f: paramWraper{[]registerID{L, A}, 4},
	0x68: paramWraper{[]registerID{L, B}, 4},
	0x69: paramWraper{[]registerID{L, C}, 4},
	0x6a: paramWraper{[]registerID{L, D}, 4},
	0x6b: paramWraper{[]registerID{L, E}, 4},
	0x6c: paramWraper{[]registerID{L, H}, 4},
	0x6d: paramWraper{[]registerID{L, L}, 4},
	0x6e: paramWraper{[]registerID{L, HL}, 8},

	0x70: paramWraper{[]registerID{HL, B}, 8},
	0x71: paramWraper{[]registerID{HL, C}, 8},
	0x72: paramWraper{[]registerID{HL, D}, 8},
	0x73: paramWraper{[]registerID{HL, E}, 8},
	0x74: paramWraper{[]registerID{HL, H}, 8},
	0x75: paramWraper{[]registerID{HL, L}, 8},
	0x36: paramWraper{[]registerID{HL, N}, 12},

	0x0a: paramWraper{[]registerID{A, BC}, 8},
	0x1a: paramWraper{[]registerID{A, DE}, 8},
	0xfa: paramWraper{[]registerID{A, NN}, 16},
	0x3e: paramWraper{[]registerID{A, N}, 8},

	0x02: paramWraper{[]registerID{BC, A}, 8},
	0x12: paramWraper{[]registerID{DE, A}, 8},
	0x77: paramWraper{[]registerID{HL, A}, 8},
	0xea: paramWraper{[]registerID{NN, A}, 16},
}

var map8bitLDAorC = map[byte]paramWraper{
	0xf2: paramWraper{[]registerID{A, C}, 8},
	0xe2: paramWraper{[]registerID{C, A}, 8},
}

// LD B,n  06 8
func (gb *GBCpu) instruction(opcode byte) (cycle int) {

	if v, ok := map8bitsLDr1r2[opcode]; ok {
		return gb._8bitsLDr1r2(v)
	}

	if v, ok := map8bitLDAorC[opcode]; ok {
		return gb._8bitsLDAorC(v)
	}
	panic(fmt.Errorf("unknown opcode %x", opcode))
}

func (gb *GBCpu) _8bitsLDr1r2(param paramWraper) (cycle int) {
	panic("to do")
	//return param.cycles
}

func (gb *GBCpu) _8bitsLDAorC(param paramWraper) (cycle int) {
	panic("to do")
}
