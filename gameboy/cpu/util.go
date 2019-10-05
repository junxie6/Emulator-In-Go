package cpu

import (
	"fmt"
	"log"
	"os"
)

func ParseInstructionfile(file os.File) []InstructionInfo {
	return nil
}

func ParseRawInstructionMap(list map[string]map[byte]paramWraper) []*InstructionInfo {
	rtn := make([]*InstructionInfo, 0x100)

	mapRegName := map[registerID]string{
		A: "A", B: "B", C: "C", D: "D",
		E: "E", F: "F", AF: "AF", BC: "BC",
		DE: "DE", H: "H", L: "L", HL: "HL",
		NN: "NN", N: "N", SP: "SP", PC: "PC",
	}

	nameformatter := map[string]string{
		"16-bits ADD": "ADD",
		"16-bits DEC": "DEC",
		"16-bits LD":  "LD",
		"16-bits INC": "INC",
		"LDAC":        "LD",
		"JPHL":        "JP",
		"JPnn":        "JP",
		"JRccn":       "JR",
		"CALLsnn":     "CALL",
	}

	for name := range list {
		for opcode := range list[name] {
			if rtn[opcode] != nil {
				log.Panic(fmt.Sprintf("duplicate opcode %02x, instruction %s and %s", opcode, name, rtn[opcode].Instruction))
			}
			rtn[opcode] = &InstructionInfo{
				Instruction: func() string {
					if _, ok := nameformatter[name]; ok {
						return nameformatter[name]
					}
					return name
				}(),
				Params: func() []string {
					p := []string{}
					for _, id := range list[name][opcode].params {
						if _, ok := mapRegName[id]; !ok {
							log.Panic("unknow id %id", id)
						}
						p = append(p, mapRegName[id])
					}
					return p
				}(),
				Opcode: fmt.Sprintf("0x%02x", opcode),
				Code:   opcode,
				Cycles: list[name][opcode].cycles,
			}
		}
	}
	return rtn
}

func atoi(s string) registerID {
	mapRegName := map[string]registerID{
		"A": A, "B": B, "C": C, "D": D,
		"E": E, "F": F, "AF": AF, "BC": BC,
		"DE": DE, "H": H, "L": L, "HL": HL,
		"NN": NN, "N": N, "SP": SP, "PC": PC,
	}
	return mapRegName[s]
}
