package main

import "fmt"

// use pointers if your struct is huge
// use pointers if you want to mutate your variable
// use pointers if you want to share your variable accross functions

// don't use pointers for things that are pointers under the hood, like slices

func changeByVal(n int) {
	n += 1
}

func changeByPtr(n *int) {
	*n += 1
}

func exercise1() {
	n := 10

	changeByVal(n)
	fmt.Println(n)

	changeByPtr(&n)
	fmt.Println(n)
}

type SomeStruct struct {
	Name string
	Age  int
}

// receiver argument
func (s SomeStruct) changeByVal() {
	s.Age = 100
}

// receiver argument
func (s *SomeStruct) changeByPtr() {
	// (*s).Age = 100

	// pointer indirection
	s.Age = 100
}

func exercise2() {
	s := SomeStruct{
		Age: 10,
	}

	s.changeByVal()
	fmt.Println(s)
	s.changeByPtr()
	fmt.Println(s)

}

func main() {
	exercise1()
	exercise2()
}
