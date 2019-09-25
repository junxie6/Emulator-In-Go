package main

import (
	"os"

	"github.com/KeyOneLi/Emulator-In-Go/chip8/cpu"
	"github.com/veandco/go-sdl2/sdl"
)

var keyValue = map[sdl.Keycode]byte{
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

func main() {
	rom := os.Args[1]
	file, err := os.Open(rom)
	must(err)

	must(sdl.Init(sdl.INIT_EVERYTHING))
	window, _ := sdl.CreateWindow("chip-8 emulator", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		cpu.ScreenW*8, cpu.ScreenH*8, sdl.WINDOW_SHOWN)
	r, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	chip8 := cpu.NewChip8(file, &Render{r})
	file.Close()
	window.UpdateSurface()

	go chip8.Run()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				code := t.Keysym.Sym
				press := t.State == 1
				if v, ok := keyValue[code]; ok {
					// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					// 	t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
					if press {
						chip8.PressKey(v)
					} else {
						chip8.ReleaseKey(v)
					}
				}
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				chip8.Close()
				break
			}
		}
	}
}
func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Render struct {
	render *sdl.Renderer
}

func (r *Render) Clear() {
	must(r.render.SetDrawColor(0, 0, 0, 0xff))
	r.render.Clear()
}

func (r *Render) Update(board [][]bool) {
	// fmt.Println("draw")
	r.Clear()
	must(r.render.SetDrawColor(0xff, 0xff, 0xff, 0xff))
	for x := range board {
		for y, v := range board[x] {
			if v {
				for i := 0; i < 8; i++ {
					for j := 0; j < 8; j++ {
						must(r.render.DrawPoint(int32(x*8+i), int32(y*8+j)))
					}
				}
			}

		}
	}
	r.render.Present()
}
