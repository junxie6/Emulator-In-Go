package cpu

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Chip8 struct {
	v              []byte //Register V0-F
	programCounter uint16
	vi             uint16
	memory         []byte

	stack        []uint16
	stackPointer byte

	drawFlag    bool
	display     Display
	screenBoard [][]bool

	keyMux         *sync.Mutex
	keySignal      *sync.Cond
	lastPressedKey byte
	KeyState       uint16

	delayTimer byte
	soundTimer byte

	random *rand.Rand

	done chan interface{}
}

func NewChip8(rom io.Reader, display Display) (c *Chip8) {
	defer func() {
		c.loadRom(rom)
		c.loadFont()
	}()
	return &Chip8{
		v:              make([]byte, registerSize),
		programCounter: programCounterStartPoint,
		memory:         make([]byte, memorySize),

		stack:   make([]uint16, stackSize),
		display: display,
		screenBoard: func() [][]bool {
			rtn := make([][]bool, ScreenW*8)
			for i := range rtn {
				rtn[i] = make([]bool, ScreenH*8)
			}
			return rtn
		}(),

		delayTimer: timer,
		soundTimer: timer,

		keySignal: sync.NewCond(&sync.Mutex{}),
		keyMux:    &sync.Mutex{},

		random: rand.New(rand.NewSource(time.Now().Unix())),
		done:   make(chan interface{}),
	}
}

func (c *Chip8) Run() {
	cycle := 60
	for {
		select {
		case <-c.done:
			return
		case <-time.After(time.Second / 1000):
			opcode := c.loadOpcode()
			c.instruction(opcode)

			c.timerUpdate()

			cycle--
			if cycle == 0 {
				c.show()
				cycle = 60
			}
		}

	}
}
func (c *Chip8) Close() {
	close(c.done)
}

func (c *Chip8) loadRom(bf io.Reader) {
	_, err := bf.Read(c.memory[programCounterStartPoint:])
	must(err)
}
func (c *Chip8) loadFont() {
	for i := range fontset {
		c.memory[i] = fontset[i]
	}
}
func (c *Chip8) timerUpdate() {
	c.delayTimer--
	c.soundTimer--

	if c.delayTimer < 0 {
		c.delayTimer = timer

	}
	if c.soundTimer < 0 {
		c.soundTimer = timer
	}
}

func (c *Chip8) loadOpcode() uint16 {
	hi := uint16(c.readMemory(c.programCounter)) << 8
	low := uint16(c.readMemory(c.programCounter + 1))
	c.skip()
	return hi | low
}

func (c *Chip8) instruction(opcode uint16) {

	switch opcode & 0xf000 {
	case 0x0000:
		if opcode == 0x00e0 {
			c.screenBoard = make([][]bool, ScreenW)
			for i := range c.screenBoard {
				c.screenBoard[i] = make([]bool, ScreenH)
			}
			c.display.Clear()
		} else if opcode == 0x00ee {
			c.stackPointer--
			if c.stackPointer < 0 {
				panic("stack overflow")
			}
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

func (c *Chip8) PressKey(key byte) {
	c.keyMux.Lock()
	defer c.keyMux.Unlock()

	c.keySignal.L.Lock()
	c.KeyState |= 1 << key
	c.lastPressedKey = key
	c.keySignal.Signal()
	c.keySignal.L.Unlock()
}

func (c *Chip8) ReleaseKey(key byte) {
	c.keyMux.Lock()
	defer c.keyMux.Unlock()

	c.KeyState &= ^(1 << key)
}

func (c *Chip8) keyPressed(key byte) bool {
	c.keyMux.Lock()
	defer c.keyMux.Unlock()

	k := uint16(1 << key)
	return c.KeyState&k != 0
}

func (c *Chip8) getKey() byte {
	c.keySignal.L.Lock()
	c.keySignal.Wait()
	c.keySignal.L.Unlock()

	return c.lastPressedKey
}

func (c *Chip8) draw(x, y, n byte) {
	c.drawFlag = true

	c.v[0xf] = 0
	for i := byte(0); i < n; i++ {
		px := c.readMemory(c.vi + uint16(i))
		for j := byte(0); j < 8; j++ {
			flag := (px>>uint(7-j))&0x1 == 1
			pre := c.screenBoard[x+j][y+i]
			if flag != pre {
				c.screenBoard[x+j][y+i] = true
			} else {
				c.screenBoard[x+j][y+i] = false
				if pre {
					c.v[0xf] = 1
				}
			}
		}
	}
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

func (c *Chip8) show() {
	if c.drawFlag {
		c.display.Update(c.screenBoard)
		c.drawFlag = false
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
