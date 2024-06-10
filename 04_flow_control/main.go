package main

import (
	"fmt"
	"strings"
)

// write a slice with elements 100, 45, 67, 86, 12 and compute their sum.
func exercise1() {

	s := 0
	arr := []int{100, 45, 67, 86, 12}
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}

	for idx, value := range arr {
		fmt.Println(idx, value)
	}

	fmt.Println(s)

}

// Alex, Andrei, Alexandra, Daniel, Iulia, Armand,  Alex, Calin
// strings
// "dr"
func exercise2() {
	arr := []string{"Alex", "Andrei", "Alexandra", "Daniel", "Iulia", "Armand", "Alex", "Calin"}

	for _, name := range arr {
		if strings.Contains(name, "dr") {
			fmt.Println(name)
		}
	}
}

// Alex, Andrei, Alexandra, Daniel, Iulia, Armand,  Alex, Calin
// strings
// "dr"
// create a frequency map with the names
func exercise3() {
	arr := []string{"Alex", "Andrei", "Alexandra", "Daniel", "Iulia", "Armand", "Alex", "Calin"}

	var names = make(map[string]int)

	x, y := names["hello"]
	fmt.Println(x, y)

	for _, name := range arr {

		// check if something is in map
		if _, ok := names[name]; ok {
			names[name]++
		} else {
			names[name] = 1
		}

		// names[name]++
	}

	fmt.Println(names)

	for name, freq := range names {
		fmt.Println(name, freq)
	}
}

func exercise4() {
	var cond bool
	cond = true
	i := 0
	// no while in Go
	for cond {
		i++
		if i == 5 {
			cond = false
		}
		fmt.Println(i)
	}
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}
