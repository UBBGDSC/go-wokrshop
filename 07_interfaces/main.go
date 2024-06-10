package main

import "fmt"

type HelloWordable interface {
	HelloWorld()
}

type RomanianHelloWorder struct {
	name string
}

func (r RomanianHelloWorder) HelloWorld() {
	fmt.Println("Hei ba!")
}

func (r RomanianHelloWorder) CeFaci() {
	fmt.Println("ce faci")
}

type EnglishHelloWorlder struct {
}

func (e EnglishHelloWorlder) HelloWorld() {
	fmt.Println("Hey mate!")
}

func main() {
	var sayshello HelloWordable

	sayshello = RomanianHelloWorder{
		name: "alex",
	}
	sayshello.HelloWorld()

	// not possible!!!
	// sayshello.CeFaci()

	// type assertion!
	romanian, ok := sayshello.(RomanianHelloWorder)
	fmt.Println(romanian, ok)
	romanian.CeFaci()

	english, ok := sayshello.(EnglishHelloWorlder)
	fmt.Println(english, ok)

	sayshello = EnglishHelloWorlder{}
	sayshello.HelloWorld()

	switch sayshello.(type) {
	case EnglishHelloWorlder:
		fmt.Println("e english")
	case RomanianHelloWorder:
		fmt.Println("romanian")
	}

	// compile time check
	var _ HelloWordable = RomanianHelloWorder{}
}
