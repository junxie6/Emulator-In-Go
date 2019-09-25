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

	delayTimer byte
	soundTimer byte

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
		nn := opcode & 0x00ff
		x := (opcode >> 8) & 0x0f

		if c.v[x] == byte(nn) {
			c.skip()
		}

	case 0x4000:
		nn := opcode & 0x00ff
		x := (opcode >> 8) & 0xf

		if c.v[x] != byte(nn) {
			c.skip()
		}

	case 0x5000:
		x := (opcode >> 8) & 0xf
		y := (opcode >> 4) & 0xf
		if c.v[x] == c.v[y] {
			c.skip()
		}
	case 0x6000:
		x := (opcode >> 8) & 0x0f
		nn := opcode & 0x00ff
		c.v[x] = byte(nn)
	case 0x7000:
		x := (opcode >> 8) & 0x0f
		nn := opcode & 0x00ff

		c.v[x] += byte(nn)
	case 0x8000:
		last := opcode & 0x000f
		x := (opcode >> 8) & 0x0f
		y := (opcode >> 4) & 0x0f
		switch last {
		case 0:
			c.v[x] = c.v[y]
		case 1:
			c.v[x] |= c.v[y]
		case 2:
			c.v[x] &= c.v[y]
		case 3:
			c.v[x] ^= c.v[y]
		case 4:
			sum := uint16(c.v[x]) + uint16(c.v[y])
			if sum > math.MaxUint8 {
				c.v[0xf] = 1
			} else {
				c.v[0xf] = 0
			}
			c.v[x] = byte(sum & 0xff)
		case 5:
			if c.v[x] > c.v[y] {
				c.v[0xf] = 1
			} else {
				c.v[0xf] = 0
			}
			c.v[x] -= c.v[y]
		case 6:
			c.v[0xf] = c.v[x] & 0x1
			c.v[x] >>= 1
		case 7:
			if c.v[y] > c.v[x] {
				c.v[0xf] = 1
			} else {
				c.v[0xf] = 0
			}
			c.v[x] = c.v[y] - c.v[x]
		case 0xe:
			c.v[0xf] = c.v[x] >> 7
			c.v[x] <<= 1
		default:
			panic(fmt.Errorf("uknown opcode %x", opcode))
		}
	case 0x9000:
		x := (opcode >> 8) & 0xf
		y := (opcode >> 4) & 0x0f
		if c.v[x] != c.v[y] {
			c.skip()
		}
	case 0xa000:
		nnn := opcode & 0x0fff
		c.vi = nnn
	case 0xb000:
		nnn := opcode & 0x0fff
		c.programCounter = nnn + uint16(c.v[0])
	case 0xc000:
		x := (opcode >> 8) & 0xf
		nn := byte(opcode & 0x00ff)
		c.v[x] = byte(c.random.Intn(math.MaxUint8)) & nn

	case 0xd000:
		n := byte(opcode & 0x0f)
		x := (opcode >> 8) & 0x0f
		y := (opcode >> 4) & 0x0f
		c.draw(c.v[x], c.v[y], n)

	case 0xe000:
		x := (opcode >> 8) & 0x0f
		nn := opcode & 0xff
		switch nn {
		case 0x9e:
			if c.keyPressed(c.v[x]) {
				c.skip()
			}
		case 0xa1:
			if !c.keyPressed(c.v[x]) {
				c.skip()
			}
		default:
			panic(fmt.Errorf("uknown opcode %x", opcode))
		}
	case 0xf000:
		x := (opcode >> 8) & 0x0f
		nn := opcode & 0xff

		switch nn {
		case 0x07:
			c.v[x] = c.delayTimer
		case 0x0a:
			c.v[x] = c.getKey()
		case 0x15:
			c.delayTimer = c.v[x]
		case 0x18:
			c.soundTimer = c.v[x]
		case 0x1e:
			c.vi += uint16(c.v[x])
		case 0x29:
			c.vi = c.builtInSpriteAddr(c.v[x])
		case 0x33:
			j, k, l := BCD(c.v[x])
			c.writeMemory(c.vi, j)
			c.writeMemory(c.vi, k)
			c.writeMemory(c.vi, l)
		case 0x55:
			for i := uint16(0); i <= x; i++ {
				c.writeMemory(i+c.vi, c.v[i])
			}
		case 0x65:
			for i := uint16(0); i <= x; i++ {
				c.v[i] = c.readMemory(i + c.vi)
			}
		default:
			panic(fmt.Sprintf("unknow opcode %x", opcode))
		}
	default:
		panic(fmt.Sprintf("unknow opcode %x", opcode))
	}
}

func (c *Chip8) keyPressed(key byte) bool {
	c.keyMux.Lock()
	defer c.keyMux.Unlock()

	k := uint16(1 << key)
	return c.KeyState&k == k
}

func (c *Chip8) getKey() byte {
	panic("getKey needed")
	return 0
}

func (c *Chip8) draw(x, y, n byte) {
	panic("need draw")
}

func (c *Chip8) builtInSpriteAddr(n byte) uint16 {
	return uint16(n) * 5
}

func (c *Chip8) skip() {
	c.programCounter += 2
}

func (c *Chip8) writeMemory(addr uint16, value byte) {
	c.memory[addr] = value
}
func (c *Chip8) readMemory(addr uint16) byte {
	return c.memory[addr]
}
