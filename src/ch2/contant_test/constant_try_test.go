package contant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstTry(t *testing.T) {
	fmt.Println("before fail")
	t.Log(Monday, Tuesday)
}

func TestConstTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
