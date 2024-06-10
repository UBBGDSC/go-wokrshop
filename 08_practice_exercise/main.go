package main

import "slices"

/*
- Adding a new programmer
- Listing all programmers (as an array)
- Send the top programmer to work (based on language and yoe constraints)
-- programmer needs to work in that language
-- programmer needs to have at least that many yoe

we want 2 implementations
*/

type Programmer struct {
	Name     string
	Language string
	YoE      int
}

type BenchPool interface {
	AddProgrammer(p Programmer)
	ListAllProgrammers() []Programmer

	// SendToWork(language string, minYears int) Programmer
}

type ArrayBenchPool struct {
	programmers []Programmer
}

func (pool *ArrayBenchPool) AddProgrammer(p Programmer) {
	pool.programmers = append(pool.programmers, p)
}

func (pool *ArrayBenchPool) ListAllProgrammers() []Programmer {
	return slices.Clone(pool.programmers)
}

func main() {

}
