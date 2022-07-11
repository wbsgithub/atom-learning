package map_ext_test

import (
	"strings"
	"testing"
)

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

func TestWcFunc(t *testing.T) {
	wc := wcFunc("I love my work and I love my family too")
	t.Log(wc)
}

func wcFunc(data string) map[string]int {
	wc := make(map[string]int)
	fields := strings.Fields(data)
	for _, key := range fields {
		if v, ok := wc[key]; ok {
			wc[key] = v + 1
		} else {
			wc[key] = 1
		}
	}
	return wc
}
