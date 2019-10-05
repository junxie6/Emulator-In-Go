package cpu

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestParseRawInstructionMap(t *testing.T) {
	target := "./instruction"
	os.Remove(target)
	outputfile, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer outputfile.Close()
	list := map[string]map[byte]paramWraper{
		"16-bits ADD": map16bitADD,
		"16-bits DEC": map16bitDEC,
		"16-bits LD":  map16bitsLD,
		"16-bits INC": map16bitINC,
		"STOP":        mapSTOP,
		"ADC":         map8bitADC,
		"ADD":         map8bitADD,
		"AND":         map8bitAND,
		"CP":          map8bitCP,
		"DEC":         map8bitDEC,
		"INC":         map8bitINC,
		"LDAC":        map8bitLDAC,
		"OR":          map8bitOR,
		"SBC":         map8bitSBC,
		"SUB":         map8bitSUB,
		"XOR":         map8bitXOR,
		"LD":          map8bitsLD,
		"LDD":         map8bitsLDD,
		"LDH":         map8bitsLDH,
		"LDI":         map8bitsLDI,
		"CCF":         mapCCF,
		"CPL":         mapCPL,
		"CALL":        mapCallsccnn,
		"CALLsnn":     mapCallsnn,
		"DAA":         mapDAA,
		"DI":          mapDI,
		"EI":          mapEI,
		"HALT":        mapHALT,
		"JPHL":        mapJPHL,
		"JP":          mapJPccnn,
		"JPnn":        mapJPnn,
		"JRccn":       mapJRccn,
		"JRn":         mapJRn,
		"NOP":         mapNOP,
		"POP":         mapPOP,
		"PUSH":        mapPUSH,
		"RET":         mapRET,
		"RETI":        mapRETI,
		"RLA":         mapRLA,
		"RLCA":        mapRLCA,
		"RRA":         mapRRA,
		"RRCA":        mapRRCA,
		"RST":         mapRST,
		"SCF":         mapSCF,
	}

	instructions := ParseRawInstructionMap(list)
	encoder := json.NewEncoder(outputfile)
	log.Println(len(instructions))
	for _, ins := range instructions {
		if err := encoder.Encode(&ins); err != nil {
			panic(err)
		}
	}
}
