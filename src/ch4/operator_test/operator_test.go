package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	t.Log(Readable)
	a = a &^ Readable //0001 ->   0110
	t.Log(Executable)
	a = a &^ Executable // 0100  -> 0011
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1,2,3,5,6}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a==c)
	t.Log(a == d)
}
