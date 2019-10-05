package cpu

import "fmt"

var mapReg_ld_nn_n8bits = map[byte]registerID{}

type paramWraper struct {
	params []registerID
	cycles int
}

var map8bitsLD = map[byte]paramWraper{
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

var map8bitsLDD = map[byte]paramWraper{
	0x3a: paramWraper{[]registerID{A, HL}, 8},
	0x32: paramWraper{[]registerID{HL, A}, 8},
}

var map8bitsLDI = map[byte]paramWraper{
	0x2a: paramWraper{[]registerID{A, HL}, 8},
	0x22: paramWraper{[]registerID{HL, A}, 8},
}

var map8bitsLDH = map[byte]paramWraper{
	0xe0: paramWraper{[]registerID{N, A}, 12},
	0xf0: paramWraper{[]registerID{A, N}, 12},
}
var map8bitLDAC = map[byte]paramWraper{
	0xf2: paramWraper{[]registerID{A, C}, 8},
	0xe2: paramWraper{[]registerID{C, A}, 8},
}
var map16bitsLD = map[byte]paramWraper{
	0x01: paramWraper{[]registerID{BC, NN}, 12},
	0x11: paramWraper{[]registerID{DE, NN}, 12},
	0x21: paramWraper{[]registerID{HL, NN}, 12},
	0x31: paramWraper{[]registerID{HL, NN}, 12},

	0xf9: paramWraper{[]registerID{SP, HL}, 8},
	0xf8: paramWraper{[]registerID{SP, N}, 12},

	0x08: paramWraper{[]registerID{NN, SP}, 20},
}

var mapPUSH = map[byte]paramWraper{
	0xf5: paramWraper{[]registerID{AF}, 16},
	0xc5: paramWraper{[]registerID{BC}, 16},
	0xd5: paramWraper{[]registerID{DE}, 16},
	0xe5: paramWraper{[]registerID{HL}, 16},
}

var mapPOP = map[byte]paramWraper{
	0xf1: paramWraper{[]registerID{AF}, 12},
	0xc1: paramWraper{[]registerID{BC}, 12},
	0xd1: paramWraper{[]registerID{DE}, 12},
	0xe1: paramWraper{[]registerID{HL}, 12},
}

// 8bit ALU
var map8bitADD = map[byte]paramWraper{
	0x87: paramWraper{[]registerID{A, A}, 4},
	0x80: paramWraper{[]registerID{A, B}, 4},
	0x81: paramWraper{[]registerID{A, C}, 4},
	0x82: paramWraper{[]registerID{A, D}, 4},
	0x83: paramWraper{[]registerID{A, E}, 4},
	0x84: paramWraper{[]registerID{A, H}, 4},
	0x85: paramWraper{[]registerID{A, L}, 4},
	0x86: paramWraper{[]registerID{A, HL}, 4},
	0xc6: paramWraper{[]registerID{A, N}, 8},
}

var map8bitADC = map[byte]paramWraper{
	0x8f: paramWraper{[]registerID{A, A}, 4},
	0x88: paramWraper{[]registerID{A, B}, 4},
	0x89: paramWraper{[]registerID{A, C}, 4},
	0x8a: paramWraper{[]registerID{A, D}, 4},
	0x8b: paramWraper{[]registerID{A, E}, 4},
	0x8c: paramWraper{[]registerID{A, F}, 4},
	0x8d: paramWraper{[]registerID{A, H}, 4},
	0x8e: paramWraper{[]registerID{A, L}, 8},
	0xce: paramWraper{[]registerID{A, HL}, 8},
}

var map8bitSUB = map[byte]paramWraper{
	0x97: paramWraper{[]registerID{A}, 4},
	0x90: paramWraper{[]registerID{B}, 4},
	0x91: paramWraper{[]registerID{C}, 4},
	0x92: paramWraper{[]registerID{D}, 4},
	0x93: paramWraper{[]registerID{E}, 4},
	0x94: paramWraper{[]registerID{H}, 4},
	0x95: paramWraper{[]registerID{L}, 4},
	0x96: paramWraper{[]registerID{HL}, 8},
	0xd6: paramWraper{[]registerID{N}, 8},
}

var map8bitSBC = map[byte]paramWraper{
	0x9f: paramWraper{[]registerID{A, A}, 4},
	0x98: paramWraper{[]registerID{A, B}, 4},
	0x99: paramWraper{[]registerID{A, B}, 4},
	0x9a: paramWraper{[]registerID{A, B}, 4},
	0x9b: paramWraper{[]registerID{A, B}, 4},
	0x9c: paramWraper{[]registerID{A, B}, 4},
	0x9d: paramWraper{[]registerID{A, B}, 4},
	0x9e: paramWraper{[]registerID{A, B}, 8},
	// 0x??: paramWraper{[]registerID{?, ?}, ?}, ????????
}

// LD B,n  06 8
func (gb *GBCpu) instruction(opcode byte) (cycle int) {

	if v, ok := map8bitsLD[opcode]; ok {
		return gb._8bitsLDr1r2(v)
	}

	if v, ok := map8bitLDAC[opcode]; ok {
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
