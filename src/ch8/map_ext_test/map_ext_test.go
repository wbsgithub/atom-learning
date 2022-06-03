package map_ext_test

import "testing"

func TestMapExt(t *testing.T) {
	m := map[int]func(p int) int{}
	m[1] = func(p int) int {
		return p
	}
	m[2] = func(p int) int {
		return p * p
	}
	m[3] = func(p int) int {
		return p * p * p
	}
	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](2))
}

func TestMapForSet(t *testing.T) {
	m := map[int]bool{}
	m[1] = true
	n := 1
	t.Log(m[2])
	if m[n] {
		t.Logf("%d is exsting", n)
	} else {
		t.Logf("%d is not exsting", n)
	}
	t.Logf("map len is %d", len(m))
	delete(m, 1)
	n = 1
	if m[n] {
		t.Logf("%d is exsting", n)
	} else {
		t.Logf("%d is not exsting", n)
	}
	t.Logf("map len is %d", len(m))

}
