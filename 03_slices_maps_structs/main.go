package main

import "fmt"

// arrays, slices, maps
// structs

func example1() {
	var arr [5]int
	fmt.Println(arr)

	arr[1] = 10
	fmt.Println(arr)

	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}

func example2() {
	// slice
	arr := make([]int, 0, 10)

	fmt.Println(arr, len(arr), cap(arr))

	// source of a lot of pain
	arr2 := make([]int, 5, 10)
	fmt.Println(arr2, len(arr2), cap(arr2))

	// source of a lot of pain
	arr3 := make([]int, 5)
	fmt.Println(arr3, len(arr3), cap(arr3))

	arr4 := []int{1, 2, 3}
	fmt.Println(arr4, len(arr4), cap(arr4))
}

func example3() {
	arr := []int{1, 2, 3}
	fmt.Println(arr, len(arr), cap(arr))

	arr = append(arr, 4, 5)

	fmt.Println(arr, len(arr), cap(arr))

	var arr2 []int
	fmt.Println(arr2 == nil)

	arr2 = append(arr2, 10)

	var arr3 = make([]int, 0)
	fmt.Println(arr3 == nil)

}

func example4() {
	mymap := make(map[int]int, 0)

	mymap[10] = 9
	mymap[9] = 10

	fmt.Println(mymap)
	delete(mymap, 10)
	fmt.Println(mymap)
}

func example5() {
	// go idiomatic way of making a set
	set := make(map[int]struct{})
	set[10] = struct{}{}
	set[9] = struct{}{}

	fmt.Println(set)
}

type SomeStruct struct {
	Name string
	Age  int
}

func example6() {
	s1 := SomeStruct{
		Name: "Alex",
		Age:  23,
	}

	s2 := SomeStruct{}

	// format verb
	fmt.Printf("%+v, %+v \n", s1, s2)
}

type SomeOtherStruct struct {
	Job    string
	Person SomeStruct
}

type SomeOtherStructEmbed struct {
	Job string
	SomeStruct
}

func example7() {
	s1 := SomeStruct{
		Name: "Alex",
		Age:  23,
	}
	s := SomeOtherStruct{
		Job:    "Programmer",
		Person: s1,
	}

	semb := SomeOtherStructEmbed{
		Job:        "Programmer",
		SomeStruct: s1,
	}

	fmt.Println(s, semb)

	// s.Person.Age

	// semb.Age
	// semb.SomeStruct.Age
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
}
