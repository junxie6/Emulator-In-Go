package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	displayReservedStart = 0xf00
	displayReservedEnd   = 0xfff

	// 0xEA0-0xEFF
	callStackStart = 0xea0
	callStackEnd   = 0xeff

	memSize           = 0x1000
	programStartPoint = 0x200
)

var stackTail = 0
var stack = make([]uint16, 2048)

var register []byte
var memory []byte
var displayMem [][]bool
var keyPressWait *sync.Cond
var keyMux sync.Mutex
var keyStat = [16]bool{false, false, false, false,
	false, false, false, false,
	false, false, false, false,
	false, false, false, false,
}
var LastPressedKey byte

var pc uint16
var idRegister uint16

var delayTimer byte
var soundTimer byte

var random *rand.Rand

var filename string
var drawFlag bool

var chip8Fontset = []byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

func init() {
	keyPressWait = sync.NewCond(&sync.Mutex{})
	displayMem = make([][]bool, 64*8)
	for i := range displayMem {
		displayMem[i] = make([]bool, 32*8)
	}

	stackTail = 0
	filename = "Tetris"
	register = make([]byte, 16)
	memory = make([]byte, memSize)

	random = rand.New(rand.NewSource(time.Now().Unix()))

	pc = programStartPoint

	bf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(bf); i++ {
		memory[i+programStartPoint] = bf[i]
	}

	for i := 0; i < len(chip8Fontset); i++ {
		memory[i] = chip8Fontset[i]
	}
}

var window *sdl.Window
var pic *sdl.Renderer

func main() {
	must(sdl.Init(sdl.INIT_EVERYTHING))
	window, _ := sdl.CreateWindow("chip-8 emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		64*8, 32*8, sdl.WINDOW_SHOWN)

	pic, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	window.UpdateSurface()
	go run()
	running := true

	keyMap := map[sdl.Keycode]bool{
		sdl.K_1: false,
		sdl.K_2: false,
		sdl.K_3: false,
		sdl.K_4: false,
		sdl.K_q: false,
		sdl.K_w: false,
		sdl.K_e: false,
		sdl.K_r: false,
		sdl.K_a: false,
		sdl.K_s: false,
		sdl.K_d: false,
		sdl.K_f: false,
		sdl.K_z: false,
		sdl.K_x: false,
		sdl.K_c: false,
		sdl.K_v: false,
	}
	keyValue := map[sdl.Keycode]byte{
		sdl.K_1: 0x1,
		sdl.K_2: 0x2,
		sdl.K_3: 0x3,
		sdl.K_4: 0xc,
		sdl.K_q: 0x4,
		sdl.K_w: 0x5,
		sdl.K_e: 0x6,
		sdl.K_r: 0xd,
		sdl.K_a: 0x7,
		sdl.K_s: 0x8,
		sdl.K_d: 0x9,
		sdl.K_f: 0xe,
		sdl.K_z: 0xa,
		sdl.K_x: 0x0,
		sdl.K_c: 0xb,
		sdl.K_v: 0xf,
	}
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				//fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				//	t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				code := t.Keysym.Sym
				press := t.State == 1
				if _, ok := keyMap[code]; ok {
					keyMap[code] = press
					keyMux.Lock()
					keyStat[keyValue[code]] = press
					keyMux.Unlock()

					keyPressWait.L.Lock()
					if press {
						LastPressedKey = keyValue[code]
					}
					keyPressWait.Signal()
					keyPressWait.L.Unlock()
				}
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}

}

func nextOpCode() uint16 {
	x, y := memory[pc], memory[pc+1]
	fmt.Printf("pc: %x op: %02x%02x\n", pc, x, y)
	pc += 2
	return (uint16(x) << 8) | uint16(y)
}

var force = 0

func run() {
	// input := bufio.NewReader(os.Stdin)
	// counter := 60
	for {
		op := nextOpCode()
		// log.Printf("%x %d", op, pc)
		switch head(op) {
		case 0x0:
			if op == 0x00e0 {
				must(pic.SetDrawColor(0, 0, 0, 0xff))
				must(pic.Clear())
			} else if op == 0x00ee {
				pc = stack[stackTail-1]
				stackTail--
			} else {
				addr := op & 0x0fff
				pc = addr
			}
		case 0x1:
			addr := op & 0x0fff
			pc = addr
		case 0x2:
			stack[stackTail] = pc
			stackTail++
			addr := op & 0x0fff
			pc = addr
		case 0x3:
			nn := op & 0x00ff
			x := (op >> 8) & 0x0f

			if register[x] == byte(nn) {
				pc += 2
			}

		case 0x4:
			nn := op & 0x00ff
			x := (op >> 8) & 0x0f

			if register[x] != byte(nn) {
				pc += 2
			}

		case 0x5:
			x := (op >> 8) & 0x0f
			y := (op >> 4) & 0x0f
			if register[x] == register[y] {
				pc += 2
			}
		case 0x6:
			x := (op >> 8) & 0x0f
			nn := op & 0x00ff

			register[x] = byte(nn)
		case 0x7:
			x := (op >> 8) & 0x0f
			nn := op & 0x00ff

			register[x] += byte(nn)
		case 0x8:
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
		case 0x9:
			x := (op >> 8) & 0x0f
			y := (op >> 4) & 0x0f
			if register[x] != register[y] {
				pc += 2
			}
		case 0xa:
			nnn := op & 0x0fff
			idRegister = nnn
		case 0xb:
			nnn := op & 0x0fff
			pc = nnn + uint16(register[0])
		case 0xc:
			r := byte(random.Intn(256))
			x := (op >> 8) & 0x0f
			nn := byte(op & 0x00ff)
			register[x] = r & nn

		case 0xd:
			n := byte(op & 0x0f)
			x := (op >> 8) & 0x0f
			y := (op >> 4) & 0x0f
			draw(register[x], register[y], n)
			// panic("need draw")
		case 0xe:
			x := (op >> 8) & 0x0f
			nn := op & 0xff
			switch nn {
			case 0x9e:
				if key(register[x]) {
					pc += 2
					// fmt.Println("skip")
				}
				// fmt.Println("9e", register[x], key(register[x]), pc)
				// _, _, _ = input.ReadLine()
			case 0xa1:
				if !key(register[x]) {
					pc += 2
					// fmt.Println("skip")
				}
				// fmt.Printf("a1 %x %v %x\n", register[x], !key(register[x]), pc)
				// _, _, _ = input.ReadLine()

			default:
				panic(fmt.Sprintf("%x", op))
			}
		case 0xf:
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
		// if counter == 0 {
		if drawFlag {
			display()
			drawFlag = false
		}

		// _, _, _ = input.ReadLine()

	}
}

func head(op uint16) uint16 {
	return op >> 12
}

func draw(x, y, n byte) {
	// fmt.Println("draw")
	drawFlag = true
	sp := make([]byte, n)

	for i := range sp {
		sp[i] = memory[idRegister+uint16(i)]
	}

	register[0xf] = 0
	for i := byte(0); i < n; i++ {
		px := memory[idRegister+uint16(i)]

		for j := byte(0); j < 8; j++ {
			flag := (px>>uint(7-j))&0x1 == 1
			pre := displayMem[x+j][y+i]
			if flag != pre {
				displayMem[x+j][y+i] = true
			} else {
				displayMem[x+j][y+i] = false
				if pre {
					register[0xf] = 1
				}
			}
		}
	}

}
func key(k byte) bool {
	keyMux.Lock()
	defer keyMux.Unlock()

	return keyStat[k]
}
func getKey() byte {
	keyPressWait.L.Lock()
	keyPressWait.Wait()
	keyPressWait.L.Unlock()
	return LastPressedKey
}
func spriteAddr(addr byte) uint16 {
	if addr >= 16 {
		panic(addr)
	}
	// fmt.Println(addr)
	return uint16(addr) * 5
}

func bcd(data byte) (byte, byte, byte) {
	a := data / 100
	b := (data % 100) / 10
	c := data % 10

	return a, b, c
}

// 64/8 bits * 32 -> 8byte * 32

func display() {
	// fmt.Println("display")
	must(pic.SetDrawColor(0x0, 0x0, 0x0, 0xff))

	must(pic.Clear())
	must(pic.SetDrawColor(0xff, 0xff, 0xff, 0xff))

	for x := range displayMem {
		for y, v := range displayMem[x] {
			X, Y := int32(x)*8, int32(y)*8
			if v {
				for i := int32(0); i < 8; i++ {
					for j := int32(0); j < 8; j++ {
						must(pic.DrawPoint(X+i, Y+j))
					}
				}
			}
		}
	}
	pic.Present()
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}
