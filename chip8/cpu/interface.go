package cpu

type Display interface {
	Update([][]bool)
	Clear()
}
