package main

import (
	"errors"
	"fmt"
	"slices"
)

/*
Write some code with the following functionalities:

- Adding a new programmer
- Listing all programmers (as an array)
- Send the top programmer to work (based on language and yoe constraints)
-- programmer needs to work in that language
-- programmer needs to have at least that many yoe

You should ideally have two implementations. The reason will be explained later.
*/

type Programmer struct {
	// Assume names are unique, for simplicity.
	Name     string
	Language string
	YoE      int
}

type BenchPool interface {
	AddProgrammer(p Programmer)
	ListAllProgrammers() []Programmer
	SendToWork(language string, minYears int) (Programmer, error)
}

type ArrayBenchPool struct {
	programmers []Programmer
}

// idiomatic Go constructor
func NewArrayBenchPool() *ArrayBenchPool {
	return &ArrayBenchPool{
		// 16 is a completely random choice
		programmers: make([]Programmer, 0, 16),
	}
}

// Use a pointer receiver! This is going to be an object that we'll be sharing
// around since there's no point really in initiating two pools. In addition to
// that, we will be mutating the state of this object through the exposed
// operations.
//
// see the aggregates section of this article https://medium.com/stackademic/implementing-the-outbox-pattern-from-scratch-by-following-ddd-9972eae4f1ab#:~:text=Note%20that%20an%20aggregate%20is%20a%20domain%2Ddriven%20design%20concept.%20The%20main%20purpose%20of%20an%20aggregate%20is%20that%20we%E2%80%99re%20always%20internally%20protecting%20this%20object%E2%80%99s%20consistency%20through%20the%20public%20methods%20we%20expose%20to%20mutate%20it.%20Let%E2%80%99s%20also%20create%20some%20of%20these%20methods%3A
// if you want to dive more in depth.
func (pool *ArrayBenchPool) AddProgrammer(p Programmer) {
	pool.programmers = append(pool.programmers, p)
}

func (pool *ArrayBenchPool) ListAllProgrammers() []Programmer {
	// return a clone since we don't want to be able to modify the state of
	// the pool from outside our public interface.
	return slices.Clone(pool.programmers)
}

// Nice job! Time to write the tests.
func (pool *ArrayBenchPool) SendToWork(language string, minyears int) (Programmer, error) {
	var goesToWork Programmer // the return
	var location int = -1
	for idx, programmer := range pool.programmers {
		if programmer.Language != language {
			continue
		}
		if programmer.YoE >= minyears {
			goesToWork = programmer
			location = idx
			break
		}
	}
	if location == -1 {
		return Programmer{}, nil
	}
	pool.programmers = slices.Delete(pool.programmers, location, location+1)
	return goesToWork, nil
}

// Trick! Compile-time check to see if the type satisfies the interface.

var _ BenchPool = NewArrayBenchPool()

type MapBenchPool struct {
	// I prefer the name duplication, simply because it's more convenient rather
	// than defining some smaller struct excluding the name. Imagine having to
	// re-create the programmer every time in the opreations returning
	// "Programmer".
	programmers map[string][]Programmer
}

func NewMapBenchPool() *MapBenchPool {
	return &MapBenchPool{
		programmers: make(map[string][]Programmer),
	}
}

func (pool *MapBenchPool) AddProgrammer(p Programmer) {
	if _, ok := pool.programmers[p.Language]; ok {
		pool.programmers[p.Language] = append(pool.programmers[p.Language], p)
	} else {
		pool.programmers[p.Language] = []Programmer{p} // you can also initialize array like this
	}

	// The existence check is redundant and was only done to make the core easier
	// to understand. Just having the line below has the exact same effect,
	// because append works on a nil array as well.
	// pool.programmers[p.Language] = append(pool.programmers[p.Language], p)
}

func (pool *MapBenchPool) ListAllProgrammers() []Programmer {
	allProgrammers := make([]Programmer, 0)
	// loop through all languages
	for _, listOfProgrammers := range pool.programmers {
		// add all the programmers of that language
		allProgrammers = append(allProgrammers, listOfProgrammers...)
	}
	return allProgrammers
}

// Implementation remains as an exercise.
func (pool *MapBenchPool) SendToWork(language string, minYears int) (Programmer, error) {
	// this is only implemented to satisfy the interface, and is Go's more or
	// less equivalent of Python's "pass". It's not the official recommended
	// way, but I've seen this in many open source repos.
	panic("unimplemented")
}

// Sentinel error.
var ErrNoProgrammerAvailable = errors.New("no programmer is available")

// custom error type
type InvalidLanguageError struct {
	availableLanguages []string
	language           string
}

func NewInvalidLanguageError(available []string, language string) InvalidLanguageError {
	return InvalidLanguageError{
		availableLanguages: available,
		language:           language,
	}
}

// Implement the error interface.
func (err InvalidLanguageError) Error() string {
	return fmt.Sprintf("language %s is not available. we only offer %v", err.language, err.availableLanguages)
}
