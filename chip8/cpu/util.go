package cpu

func BCD(value byte) (a, b, c byte) {
	a = value / 100
	b = (value % 100) / 10
	c = value % 10

	return a, b, c
}
