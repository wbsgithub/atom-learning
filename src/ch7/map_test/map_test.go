package map_test

import "testing"

func TestMapInit(t *testing.T) {
	//第一种
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[1])
	t.Logf("m1 len %d", len(m1))
	//第二种
	//定义变量
	var m4 map[int]int
	//初始化
	m4 = make(map[int]int, 10)
	m4[2] = 2
	t.Log(m4)

	m2 := map[int]int{}
	m2[1] = 1
	t.Log(len(m2))
	t.Log(m2[2])
	//第三种 10 is cap not len
	m3 := make(map[string]int, 1)
	t.Log(len(m3))
	m3["a"] = 1
	m3["b"] = 1
	t.Log(len(m3))
	t.Log(m3["a"])
}

func TestNotExistingKey(t *testing.T) {
	m := map[string]string{}
	t.Log(m["a"])
	m["b"] = "1"
	if v, ok := m["a"]; ok {
		t.Logf("key b is %q", v)
	} else {
		t.Logf("value: %q", v)
		t.Log("key b is not existing")
	}
}
func TestTrabelMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		t.Logf("key is %s , value is %d", k, v)
	}
}
func TestTrabelMap1(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k := range m {
		t.Logf("key is %s", k)
	}
}
