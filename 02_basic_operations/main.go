package main

import (
	"fmt"
	"strconv"
)

func exercise() {
	fmt.Println("exercise")
}

// int, float64, bool, string

// 1454 + 234 - 345

func exercise1() {
	fmt.Println(1454 + 234 - 345)

	var x1, x2, x3 int
	x1 = 1454
	x2 = 234
	x3 = 345

	fmt.Println(x1 + x2 - x3)

	y1 := 1454
	y2 := 234
	y3 := 345

	fmt.Println(y1 + y2 - y3)

	var z1, z2 float64

	z1 = 19.6
	z2 = 12.3
	fmt.Println(z1 / z2)

	a1 := 10e0
	a2 := 2.5e0

	fmt.Println(a1 / a2)
}

func exercise2() {
	var x1, x2, x3 bool
	fmt.Println(x1 && x2 || !x3)
}

func exersise3() {
	alex := "alex"
	armand := "armand"

	fmt.Println(alex + armand)
	s := alex + armand

	// s[3] = 'a'
	fmt.Println(s[1:4])
}

func exercise4() {
	x := 10

	// go doc -all -src strconv.Atoi
	xstr := strconv.Itoa(x)
	fmt.Println(xstr)

	y, _ := strconv.Atoi(xstr)
	fmt.Println(y)

}

func main() {
	// exercise()
	exercise1()
	exercise2()
	exersise3()
	exercise4()
}
