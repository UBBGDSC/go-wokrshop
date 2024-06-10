package main

func division(dividend int, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor

	return quotient, remainder
}

// parametri
func divisionsh(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor

	return quotient, remainder
}

func divisionsh2(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor

	return
}

var f func(int, int) (int, int)

type dividefunc func(int, int) (int, int)

type Coordinate [2]int

func (c Coordinate) SumOfCoords() int {
	return c[1] + c[0]
}

func main() {
	// arguments
	division(7, 3)

	f = division
	var g dividefunc = division
	g(7, 3)
}
