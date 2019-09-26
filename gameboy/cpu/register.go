package cpu

//Register define 16bit register
type Register uint16

func preloadRegister() map[registerID]*Register {
	return map[registerID]*Register{
		AF: new(Register),
		BC: new(Register),
		DE: new(Register),
		HL: new(Register),
		SP: new(Register),
		PC: new(Register),
	}
}

func (r *Register) Write(value uint16) {
	*r = Register(value)
}

func (r *Register) Read() uint16 {
	return uint16(*r)
}

//WriteHi write data to hi
func (r *Register) WriteHi(hi byte) {
	*r = (Register(hi) << 8) | (*r & 0x00ff)
}

//ReadHi read data from hi
func (r *Register) ReadHi() byte {
	return byte((*r & 0xff00) >> 8)
}

//WriteLo write data to lo
func (r *Register) WriteLo(lo byte) {
	*r = (*r & 0xff00) | Register(lo)
}

//ReadLo read data from lo
func (r *Register) ReadLo() byte {
	return byte(*r & 0x00ff)
}
