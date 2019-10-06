package cpu

import "fmt"

// Register interface for register
type Register interface {
	Write(value uint16)
	Read() uint16
}

// Register8bitWrap wrap 16bit register to 8bit
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

//RegisterPool struct to easily operate Register, both 8 bits and 16 bits
type RegisterPool struct {
	list map[registerID]*Register16bit
}

//NewRegisterPool constructor for instance RegisterPool
func NewRegisterPool() *RegisterPool {
	return &RegisterPool{
		list: preloadRegister(),
	}
}

//Get get register A,F,AF,B,C,BC,D,E,DE,H,L,HL,PC,SP
func (p *RegisterPool) Get(id registerID) Register {

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

//GetFlag Z, N, H, C
func (p *RegisterPool) GetFlag(f byte) bool {
	value := byte(p.Get(F).Read())

	return value&f == f
}

//ResetFlag Z, N, H, C
func (p *RegisterPool) ResetFlag(f byte) {
	r := p.Get(F)
	value := byte(r.Read())

	value = value & (^f)
	r.Write(uint16(value))
}

//SetFlag Z, N, H, C
func (p *RegisterPool) SetFlag(f byte) {
	r := p.Get(F)
	value := byte(r.Read())
	value = value | f
	r.Write(uint16(value))
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
