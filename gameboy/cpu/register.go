package cpu

import "fmt"

type Register interface {
	Write(value uint16)
	Read() uint16
}

type Register8bitWrap struct {
	write func(value byte)
	read  func() byte
}

func (r *Register8bitWrap) Write(value uint16) {
	r.write(byte(value))
}

func (r *Register8bitWrap) Read() uint16 {
	return uint16(r.read())
}

type RegisterPool struct {
	list map[registerID]*Register16bit
}

func NewRegisterPool() *RegisterPool {
	return &RegisterPool{
		list: preloadRegister(),
	}
}

func (p *RegisterPool) Get(id registerID) Register {
	// var id registerID
	var _id registerID
	switch id {
	case A, F, AF:
		_id = AF
	case B, C, BC:
		_id = BC
	case D, E, DE:
		_id = DE
	case H, L, HL:
		_id = HL
	default:
		panic(fmt.Sprintf("unknow id %d", id))
	}

	target := p.list[_id]

	switch id {
	case AF, BC, DE, HL, SP, PC:
		return target
	case A, B, D, H:
		return &Register8bitWrap{
			write: target.WriteHi,
			read:  target.ReadLo,
		}
	case F, C, E, L:
		return &Register8bitWrap{
			write: target.WriteLo,
			read:  target.ReadLo,
		}
	default:
		panic(fmt.Sprintf("unknow id %d", id))
	}
}

//Register16bit define 16bit register
type Register16bit uint16

func preloadRegister() map[registerID]*Register16bit {
	return map[registerID]*Register16bit{
		AF: new(Register16bit),
		BC: new(Register16bit),
		DE: new(Register16bit),
		HL: new(Register16bit),
		SP: new(Register16bit),
		PC: new(Register16bit),
	}
}

func (r *Register16bit) Write(value uint16) {
	*r = Register16bit(value)
}

func (r *Register16bit) Read() uint16 {
	return uint16(*r)
}

//WriteHi write data to hi
func (r *Register16bit) WriteHi(hi byte) {
	*r = (Register16bit(hi) << 8) | (*r & 0x00ff)
}

//ReadHi read data from hi
func (r *Register16bit) ReadHi() byte {
	return byte((*r & 0xff00) >> 8)
}

//WriteLo write data to lo
func (r *Register16bit) WriteLo(lo byte) {
	*r = (*r & 0xff00) | Register16bit(lo)
}

//ReadLo read data from lo
func (r *Register16bit) ReadLo() byte {
	return byte(*r & 0x00ff)
}
