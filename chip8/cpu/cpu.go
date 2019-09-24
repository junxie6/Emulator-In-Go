package cpu

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type Chip8 struct {
	v              []byte //Register V0-F
	programCounter uint16
	vi             uint16
	memory         []byte

	stack        []uint16
	stackPointer byte

	display     Display
	screenBoard [][]bool

	keyMux         *sync.Mutex
	keySignal      *sync.Cond
	lastPressedKey byte
	KeyState       uint16

	delayTimer uint16
	soundTimer uint16

	random *rand.Rand
}

func (c *Chip8) Run() {

}

func (c *Chip8) instruction(opcode uint16) {

	switch opcode & 0xf000 {
	case 0x0000:
		if opcode == 0x00e0 {
			c.display.Clear()
		} else if opcode == 0x00ee {
			c.stackPointer--
			c.programCounter = c.stack[c.stackPointer]
		} else {
			c.programCounter = opcode & 0x0fff
		}
	case 0x1000:
		c.programCounter = opcode & 0x0fff
	case 0x2000:
		c.stack[c.stackPointer] = c.programCounter
		c.stackPointer++
		c.programCounter = opcode & 0x0fff
	case 0x3000:
		nn := op & 0x00ff
		x := (op >> 8) & 0x0f

		if register[x] == byte(nn) {
			pc += 2
		}

	case 0x4000:
		nn := op & 0x00ff
		x := (op >> 8) & 0x0f

		if register[x] != byte(nn) {
			pc += 2
		}

	case 0x5000:
		x := (op >> 8) & 0x0f
		y := (op >> 4) & 0x0f
		if register[x] == register[y] {
			pc += 2
		}
	case 0x6000:
		x := (op >> 8) & 0x0f
		nn := op & 0x00ff

		register[x] = byte(nn)
	case 0x7000:
		x := (op >> 8) & 0x0f
		nn := op & 0x00ff

		register[x] += byte(nn)
	case 0x8000:
		last := op & 0x000f
		x := (op >> 8) & 0x0f
		y := (op >> 4) & 0x0f
		switch last {
		case 0:
			register[x] = register[y]
		case 1:
			register[x] |= register[y]
		case 2:
			register[x] &= register[y]
		case 3:
			register[x] ^= register[y]
		case 4:
			tmp := uint16(register[x]) + uint16(register[y])
			if tmp > math.MaxUint8 {
				register[0xf] = 1
			} else {
				register[0xf] = 0
			}
			register[x] = byte(tmp & 0xff)
		case 5:
			// tmp := uint16(register[x]) + uint16(register[y])
			if register[x] > register[y] {
				register[0xf] = 1
			} else {
				register[0xf] = 0
			}
			register[x] -= register[y]
		case 6:
			register[0xf] = register[x] & 0x0001
			register[x] >>= 1
		case 7:
			if register[y] > register[x] {
				register[0xf] = 1
			} else {
				register[0xf] = 0
			}
			register[x] = register[y] - register[x]
		case 0xe:
			register[0xf] = register[x] >> 7
			register[x] <<= 1
		default:
			panic(fmt.Sprintf("%x", op))
		}
	case 0x9000:
		x := (op >> 8) & 0x0f
		y := (op >> 4) & 0x0f
		if register[x] != register[y] {
			pc += 2
		}
	case 0xa000:
		nnn := op & 0x0fff
		idRegister = nnn
	case 0xb000:
		nnn := op & 0x0fff
		pc = nnn + uint16(register[0])
	case 0xc000:
		r := byte(random.Intn(256))
		x := (op >> 8) & 0x0f
		nn := byte(op & 0x00ff)
		register[x] = r & nn

	case 0xd000:
		n := byte(op & 0x0f)
		x := (op >> 8) & 0x0f
		y := (op >> 4) & 0x0f
		draw(register[x], register[y], n)
		// panic("need draw")
	case 0xe000:
		x := (op >> 8) & 0x0f
		nn := op & 0xff
		switch nn {
		case 0x9e:
			if key(register[x]) {
				pc += 2
			}
		case 0xa1:
			if !key(register[x]) {
				pc += 2
			}
		default:
			panic(fmt.Sprintf("%x", op))
		}
	case 0xf000:
		x := (op >> 8) & 0x0f
		nn := op & 0xff

		switch nn {
		case 0x07:
			register[x] = delayTimer
		case 0x0a:
			register[x] = getKey()
		case 0x15:
			delayTimer = register[x]
			// if delayTimer > 60 {
			// 	delayTimer = 60
			// }
		case 0x18:
			soundTimer = register[x]
		case 0x1e:
			idRegister += uint16(register[x])
		case 0x29:
			idRegister = spriteAddr(register[x])
		case 0x33:
			v := register[x]
			b, c, d := bcd(v)
			memory[idRegister] = b
			memory[idRegister] = c
			memory[idRegister] = d
		case 0x55:
			for i := uint16(0); i <= x; i++ {
				memory[i+idRegister] = register[i]
			}
		case 0x65:
			for i := uint16(0); i <= x; i++ {
				register[i] = memory[i+idRegister]
			}
		default:
			panic(fmt.Sprintf("%x", op))
		}
	default:
		panic(fmt.Sprintf("%x", op))
	}
	delayTimer--
	if delayTimer < 0 {
		delayTimer = 60
	}
	if drawFlag {
		display()
		drawFlag = false
	}
}

func (c *Chip8) skip() {
	c.programCounter += 2
}
