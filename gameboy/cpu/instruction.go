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

var map8bitAND = map[byte]paramWraper{
	0xa7: paramWraper{[]registerID{A}, 4},
	0xa0: paramWraper{[]registerID{B}, 4},
	0xa1: paramWraper{[]registerID{C}, 4},
	0xa2: paramWraper{[]registerID{D}, 4},
	0xa3: paramWraper{[]registerID{E}, 4},
	0xa4: paramWraper{[]registerID{H}, 4},
	0xa5: paramWraper{[]registerID{L}, 4},
	0xa6: paramWraper{[]registerID{HL}, 8},
	0xe6: paramWraper{[]registerID{N}, 8},
}

var map8bitOR = map[byte]paramWraper{
	0xb7: paramWraper{[]registerID{A}, 4},
	0xb0: paramWraper{[]registerID{B}, 4},
	0xb1: paramWraper{[]registerID{C}, 4},
	0xb2: paramWraper{[]registerID{D}, 4},
	0xb3: paramWraper{[]registerID{E}, 4},
	0xb4: paramWraper{[]registerID{H}, 4},
	0xb5: paramWraper{[]registerID{L}, 4},
	0xb6: paramWraper{[]registerID{HL}, 8},
	0xf6: paramWraper{[]registerID{N}, 8},
}

var map8bitXOR = map[byte]paramWraper{
	0xaf: paramWraper{[]registerID{A}, 4},
	0xa8: paramWraper{[]registerID{B}, 4},
	0xa9: paramWraper{[]registerID{C}, 4},
	0xaa: paramWraper{[]registerID{D}, 4},
	0xab: paramWraper{[]registerID{E}, 4},
	0xac: paramWraper{[]registerID{H}, 4},
	0xad: paramWraper{[]registerID{L}, 4},
	0xae: paramWraper{[]registerID{HL}, 8},
	0xee: paramWraper{[]registerID{N}, 8},
}

var map8bitCP = map[byte]paramWraper{
	0xbf: paramWraper{[]registerID{A}, 4},
	0xb8: paramWraper{[]registerID{B}, 4},
	0xb9: paramWraper{[]registerID{C}, 4},
	0xba: paramWraper{[]registerID{D}, 4},
	0xbb: paramWraper{[]registerID{E}, 4},
	0xbc: paramWraper{[]registerID{H}, 4},
	0xbd: paramWraper{[]registerID{L}, 4},
	0xbe: paramWraper{[]registerID{HL}, 8},
	0xfe: paramWraper{[]registerID{N}, 8},
}

var map8bitINC = map[byte]paramWraper{
	0x3c: paramWraper{[]registerID{A}, 4},
	0x04: paramWraper{[]registerID{B}, 4},
	0x0c: paramWraper{[]registerID{C}, 4},
	0x14: paramWraper{[]registerID{D}, 4},
	0x1c: paramWraper{[]registerID{E}, 4},
	0x24: paramWraper{[]registerID{H}, 4},
	0x2c: paramWraper{[]registerID{L}, 4},
	0x34: paramWraper{[]registerID{HL}, 12},
}

var map8bitDEC = map[byte]paramWraper{
	0x3d: paramWraper{[]registerID{A}, 4},
	0x05: paramWraper{[]registerID{B}, 4},
	0x0d: paramWraper{[]registerID{C}, 4},
	0x15: paramWraper{[]registerID{D}, 4},
	0x1d: paramWraper{[]registerID{E}, 4},
	0x25: paramWraper{[]registerID{H}, 4},
	0x2d: paramWraper{[]registerID{L}, 4},
	0x35: paramWraper{[]registerID{HL}, 4},
}

// 16-bit Arithmetic
var map16bitADD = map[byte]paramWraper{
	0x09: paramWraper{[]registerID{HL, BC}, 4},
	0x19: paramWraper{[]registerID{HL, DE}, 4},
	0x29: paramWraper{[]registerID{HL, HL}, 4},
	0x39: paramWraper{[]registerID{HL, SP}, 4},

	0xe8: paramWraper{[]registerID{SP, N}, 16},
}

var map16bitINC = map[byte]paramWraper{
	0x03: paramWraper{[]registerID{BC}, 8},
	0x13: paramWraper{[]registerID{DE}, 8},
	0x23: paramWraper{[]registerID{HL}, 8},
	0x33: paramWraper{[]registerID{SP}, 8},
}

var map16bitDEC = map[byte]paramWraper{
	0x0b: paramWraper{[]registerID{BC}, 8},
	0x1b: paramWraper{[]registerID{DE}, 8},
	0x2b: paramWraper{[]registerID{HL}, 8},
	0x3b: paramWraper{[]registerID{SP}, 8},
}

// Miscellaneous
var mapSWAP = map[uint16]paramWraper{
	0xcb37: paramWraper{[]registerID{A}, 8},
	0xcb30: paramWraper{[]registerID{B}, 8},
	0xcb31: paramWraper{[]registerID{C}, 8},
	0xcb32: paramWraper{[]registerID{D}, 8},
	0xcb33: paramWraper{[]registerID{E}, 8},
	0xcb34: paramWraper{[]registerID{H}, 8},
	0xcb35: paramWraper{[]registerID{L}, 8},
	0xcb36: paramWraper{[]registerID{HL}, 16},
}

var mapDAA = map[byte]paramWraper{
	0x27: paramWraper{[]registerID{}, 4},
}

var mapCPL = map[byte]paramWraper{
	0x2f: paramWraper{[]registerID{}, 4},
}

var mapCCF = map[byte]paramWraper{
	0x3f: paramWraper{[]registerID{}, 4},
}

var mapSCF = map[byte]paramWraper{
	0x37: paramWraper{[]registerID{}, 4},
}

var mapNOP = map[byte]paramWraper{
	0x00: paramWraper{[]registerID{}, 4},
}

var mapHALT = map[byte]paramWraper{
	0x76: paramWraper{[]registerID{}, 4},
}

var mapSTOP = map[uint16]paramWraper{
	0x1000: paramWraper{[]registerID{}, 4},
}

var mapDI = map[byte]paramWraper{
	0xf3: paramWraper{[]registerID{}, 4},
}

var mapEI = map[byte]paramWraper{
	0xfb: paramWraper{[]registerID{}, 4},
}

// Rotates & Shifts
var mapRLCA = map[byte]paramWraper{
	0x07: paramWraper{[]registerID{}, 4},
}

var mapRLA = map[byte]paramWraper{
	0x17: paramWraper{[]registerID{}, 4},
}

var mapRRCA = map[byte]paramWraper{
	0x0f: paramWraper{[]registerID{}, 4},
}

var mapRRA = map[byte]paramWraper{
	0x1f: paramWraper{[]registerID{}, 4},
}

var mapRLC = map[uint16]paramWraper{
	0xcb07: paramWraper{[]registerID{A}, 8},
	0xcb00: paramWraper{[]registerID{B}, 8},
	0xcb01: paramWraper{[]registerID{C}, 8},
	0xcb02: paramWraper{[]registerID{D}, 8},
	0xcb03: paramWraper{[]registerID{E}, 8},
	0xcb04: paramWraper{[]registerID{H}, 8},
	0xcb05: paramWraper{[]registerID{L}, 8},
	0xcb06: paramWraper{[]registerID{HL}, 16},
}

var mapRL = map[uint16]paramWraper{
	0xcb17: paramWraper{[]registerID{A}, 8},
	0xcb10: paramWraper{[]registerID{B}, 8},
	0xcb11: paramWraper{[]registerID{C}, 8},
	0xcb12: paramWraper{[]registerID{D}, 8},
	0xcb13: paramWraper{[]registerID{E}, 8},
	0xcb14: paramWraper{[]registerID{H}, 8},
	0xcb15: paramWraper{[]registerID{L}, 8},
	0xcb16: paramWraper{[]registerID{HL}, 16},
}

var mapRRC = map[uint16]paramWraper{
	0xcb0f: paramWraper{[]registerID{A}, 8},
	0xcb08: paramWraper{[]registerID{B}, 8},
	0xcb09: paramWraper{[]registerID{C}, 8},
	0xcb0a: paramWraper{[]registerID{D}, 8},
	0xcb0b: paramWraper{[]registerID{E}, 8},
	0xcb0c: paramWraper{[]registerID{H}, 8},
	0xcb0d: paramWraper{[]registerID{L}, 8},
	0xcb0e: paramWraper{[]registerID{HL}, 16},
}

var mapRR = map[uint16]paramWraper{
	0xcb1f: paramWraper{[]registerID{A}, 8},
	0xcb18: paramWraper{[]registerID{B}, 8},
	0xcb19: paramWraper{[]registerID{C}, 8},
	0xcb1a: paramWraper{[]registerID{D}, 8},
	0xcb1b: paramWraper{[]registerID{E}, 8},
	0xcb1c: paramWraper{[]registerID{H}, 8},
	0xcb1d: paramWraper{[]registerID{L}, 8},
	0xcb1e: paramWraper{[]registerID{HL}, 16},
}

var mapSLA = map[uint16]paramWraper{
	0xcb27: paramWraper{[]registerID{A}, 8},
	0xcb20: paramWraper{[]registerID{B}, 8},
	0xcb21: paramWraper{[]registerID{C}, 8},
	0xcb22: paramWraper{[]registerID{D}, 8},
	0xcb23: paramWraper{[]registerID{E}, 8},
	0xcb24: paramWraper{[]registerID{H}, 8},
	0xcb25: paramWraper{[]registerID{L}, 8},
	0xcb26: paramWraper{[]registerID{HL}, 16},
}

var mapSRA = map[uint16]paramWraper{
	0xcb2f: paramWraper{[]registerID{A}, 8},
	0xcb28: paramWraper{[]registerID{B}, 8},
	0xcb29: paramWraper{[]registerID{C}, 8},
	0xcb2a: paramWraper{[]registerID{D}, 8},
	0xcb2b: paramWraper{[]registerID{E}, 8},
	0xcb2c: paramWraper{[]registerID{H}, 8},
	0xcb2d: paramWraper{[]registerID{L}, 8},
	0xcb2e: paramWraper{[]registerID{HL}, 16},
}

var mapSRL = map[uint16]paramWraper{
	0xcb3f: paramWraper{[]registerID{A}, 8},
	0xcb38: paramWraper{[]registerID{B}, 8},
	0xcb39: paramWraper{[]registerID{C}, 8},
	0xcb3a: paramWraper{[]registerID{D}, 8},
	0xcb3b: paramWraper{[]registerID{E}, 8},
	0xcb3c: paramWraper{[]registerID{H}, 8},
	0xcb3d: paramWraper{[]registerID{L}, 8},
	0xcb3e: paramWraper{[]registerID{HL}, 16},
}

// bit opcodes
var mapBITbr = map[uint16]paramWraper{
	0xcb47: paramWraper{[]registerID{A}, 8},
	0xcb40: paramWraper{[]registerID{B}, 8},
	0xcb41: paramWraper{[]registerID{C}, 8},
	0xcb42: paramWraper{[]registerID{D}, 8},
	0xcb43: paramWraper{[]registerID{E}, 8},
	0xcb44: paramWraper{[]registerID{H}, 8},
	0xcb45: paramWraper{[]registerID{L}, 8},
	0xcb46: paramWraper{[]registerID{HL}, 16},
}

var mapSETbr = map[uint16]paramWraper{
	0xcbc7: paramWraper{[]registerID{A}, 8},
	0xcbc0: paramWraper{[]registerID{B}, 8},
	0xcbc1: paramWraper{[]registerID{C}, 8},
	0xcbc2: paramWraper{[]registerID{D}, 8},
	0xcbc3: paramWraper{[]registerID{E}, 8},
	0xcbc4: paramWraper{[]registerID{H}, 8},
	0xcbc5: paramWraper{[]registerID{L}, 8},
	0xcbc6: paramWraper{[]registerID{HL}, 16},
}

var mapRESbr = map[uint16]paramWraper{
	0xcb87: paramWraper{[]registerID{A}, 8},
	0xcb80: paramWraper{[]registerID{B}, 8},
	0xcb81: paramWraper{[]registerID{C}, 8},
	0xcb82: paramWraper{[]registerID{D}, 8},
	0xcb83: paramWraper{[]registerID{E}, 8},
	0xcb84: paramWraper{[]registerID{H}, 8},
	0xcb85: paramWraper{[]registerID{L}, 8},
	0xcb86: paramWraper{[]registerID{HL}, 16},
}

// Jumps
var mapJPnn = map[byte]paramWraper{
	0xC3: paramWraper{[]registerID{NN}, 12},
}

var mapJPccnn = map[byte]paramWraper{
	0xc2: paramWraper{[]registerID{NN}, 12},
	0xca: paramWraper{[]registerID{NN}, 12},
	0xd2: paramWraper{[]registerID{NN}, 12},
	0xda: paramWraper{[]registerID{NN}, 12},
}

var mapJPHL = map[byte]paramWraper{
	0xe9: paramWraper{[]registerID{HL}, 4},
}

var mapJRn = map[byte]paramWraper{
	0x18: paramWraper{[]registerID{N}, 8},
}

var mapJRccn = map[byte]paramWraper{
	0x20: paramWraper{[]registerID{N}, 8},
	0x28: paramWraper{[]registerID{N}, 8},
	0x30: paramWraper{[]registerID{N}, 8},
	0x38: paramWraper{[]registerID{N}, 8},
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
