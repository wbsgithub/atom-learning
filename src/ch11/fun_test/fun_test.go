package fun_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestReturnMultiValues(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	spent := timeSpent(slowFun)
	t.Log(spent(20))

	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))

}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear Resources")
	}()
	t.Log("Started")
	panic("Fatal Error")
	fmt.Println("asdfs")
}

func Clear() {
	fmt.Println("Clear Resources")
}
