package cpu

import (
	"fmt"
)

var mapReg_ld_nn_n8bits = map[byte]registerID{}

type paramWraper struct {
	params []registerID
	cycles int
}

// InstructionInfo shows detail info of an instruction
type InstructionInfo struct {
	Opcode      string   `json:"opcode"`
	Instruction string   `json:"instruction"`
	Params      []string `json:"params"`
	Code        byte     `json:"code"`
	Cycles      int      `json:"cycles"`
}

// LD B,n  06 8
func (gb *GBCpu) instructions(opcode byte) (cycle int) {

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
