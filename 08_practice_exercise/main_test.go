package main

import (
	"errors"
	"testing"
)

// first version of the test
func Test_AddProgrammer(t *testing.T) {
	p := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}

	// create a pool and add the programmer
	pool := NewArrayBenchPool()
	pool.AddProgrammer(p)

	// let's get the programmer
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 1 {
		// the other test would panic if the programmer is not there, so
		// stop here. Fatalf stops the execution of the current test.
		t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
	}

	actualProgrammer := programmers[0]
	if p != actualProgrammer {
		// In contrast, Errorf does not stop the execution.
		t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
	}
}

/*
I want you to notice that we dind't really make use of the actual Pool type, so
it didn't really matter if we were using an Array or Map bench pool. This is
because we only used exposed methods that were part of the interface. This is
generally nice, because in this unit test, we don't particualrly care about
implementation details. We want to make sure that the public methods are working
as expected.

This is however, not always possible, as there are certain cases where you
can't really test this way. For example, if we didn't expose a List method,
we'd need to define some custom method for testing that a programmer was indeed
added. This is a situation I ran across while working and is something to keep
in mind.

But, since we only used the interface methods, we can also do the following.
*/

// second version of the test. everything is exactly the same except the pool
// type

func TestAddProgrammerInterface(t *testing.T) {
	p := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}

	// create a pool and add the programmer, but the pool should be of type
	// interface. you will see how this is useful in the next version of the
	// test
	var pool BenchPool
	pool = NewArrayBenchPool()

	// Notice that changing the type is only a matter of updating the constructor
	// now, and the test is still valid for the new type!
	// pool = NewMapBenchPool()
	pool.AddProgrammer(p)

	// let's get the programmer
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 1 {
		// the other test would panic if the programmer is not there, so
		// stop here
		t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
	}

	actualProgrammer := programmers[0]
	if p != actualProgrammer {
		t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
	}
}

/*
Since you saw that we can use the exact same test for multiple implementations
of the same interface, let's do what's known as a table driven test.

https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
*/

// This test naming convention is not specific to Go. Loosely inspired from
// https://www.softwaretestingmagazine.com/knowledge/how-to-choose-the-right-name-for-unit-tests/
func Test_AddProgrammer_ProgrammerIsAdded(t *testing.T) {

	// Define the data taht each tests holds.
	type testcase struct {
		pool BenchPool
	}
	testcases := []testcase{
		{
			pool: NewArrayBenchPool(),
		},
		{
			pool: NewMapBenchPool(),
		},
	}
	// And now let's test each individual pool.
	for _, tt := range testcases {

		p := Programmer{
			Name:     "Alex",
			Language: "Go",
			YoE:      10, // definitely not lying on the resume here
		}
		tt.pool.AddProgrammer(p)

		programmers := tt.pool.ListAllProgrammers()

		if len(programmers) != 1 {
			t.Fatalf("expected %d programmers, got %d\n", 1, len(programmers))
		}

		actualProgrammer := programmers[0]
		if p != actualProgrammer {
			t.Errorf("expected to contain %v, got %v\n", p, actualProgrammer)
		}
	}
}

// Run this test only
// go test -run=HasProgrammer 8.go 8_test.go
// go help test
// go help testflag
func Test_SendToWork_HasProgrammer_ThatProgrammerIsSentToWork(t *testing.T) {
	alex := Programmer{
		Name:     "Alex",
		Language: "Go",
		YoE:      10, // definitely not lying on the resume here
	}
	armand := Programmer{
		Name:     "Armand",
		Language: "Go",
		YoE:      1,
	}
	// ask from the crowd
	ioana := Programmer{
		Name:     "Ioana",
		Language: "Java", // :(
		YoE:      2,      // definitely not lying on the resume here
	}

	var pool BenchPool
	pool = NewArrayBenchPool()

	// Now we're adding all programmers
	pool.AddProgrammer(armand) // armand first
	pool.AddProgrammer(alex)
	pool.AddProgrammer(ioana)

	// let's keep the original test, just to have an extra guard
	programmers := pool.ListAllProgrammers()

	if len(programmers) != 3 {
		t.Fatalf("expected %d programmers, got %d\n", 3, len(programmers))
	}

	// expecting to send alex
	working, _ := pool.SendToWork("Go", 5)
	if working != alex {
		t.Errorf("expected to send %v, got %v\n", alex, working)
	}

	// the programmer should be out of the pool
	programmers = pool.ListAllProgrammers()
	if len(programmers) != 2 {
		t.Fatalf("expected %d programmers, got %d\n", 2, len(programmers))
	}
}

// Now, you may notice that testing is a bit awkward for the following tests.

func Test_SendToWork_NotEnoughYoe_NoProgrammerIsSentToWork(t *testing.T) {
	// Let's remove alex from the test.

	armand := Programmer{
		Name:     "Armand",
		Language: "Go",
		YoE:      1,
	}
	ioana := Programmer{
		Name:     "Ioana",
		Language: "Java",
		YoE:      2,
	}

	var pool BenchPool
	pool = NewArrayBenchPool()

	// Remove alex
	pool.AddProgrammer(armand)
	pool.AddProgrammer(ioana)

	// So now let's send a Go programmer with 5 years of experience. But
	// we don't have one:(
	_, err := pool.SendToWork("Go", 5)
	if !errors.Is(err, ErrNoProgrammerAvailable) {
		t.Errorf("expected error %v, got %v instead\n", ErrNoProgrammerAvailable, err)
	}
}

// Let's write a quick additional test. We want an error if there is no
// language.
func Test_SendToWork_NoExistingLanguage_NoProgrammerIsSentToWork(t *testing.T) {
	// Let's remove alex from the test.

	armand := Programmer{
		Name:     "Armand",
		Language: "Go",
		YoE:      1,
	}
	ioana := Programmer{
		Name:     "Ioana",
		Language: "Java",
		YoE:      2,
	}

	var pool BenchPool
	pool = NewArrayBenchPool()

	pool.AddProgrammer(armand)
	pool.AddProgrammer(ioana)

	_, err := pool.SendToWork(".NET", 5)

	// Now that we have the new language error, we can use a different
	// way of checking errors.
	var langerr InvalidLanguageError
	if !errors.As(err, &langerr) {
		t.Errorf("expected language error %v, got %v instead\n", ErrNoProgrammerAvailable, err)
	}
}
