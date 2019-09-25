package cpu

func BCD(value byte) (a, b, c byte) {
	a := data / 100
	b := (data % 100) / 10
	c := data % 10

	return a, b, c
}
