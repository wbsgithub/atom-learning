package slice_test

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	//定义一个slice  []没有写明长度
	//var s0 []int
	//t.Log(len(s0), cap(s0))
	//s0 = append(s0, 1)
	//t.Log(len(s0), cap(s0))

	//s1 := []int{1, 2, 3, 4}
	//t.Log(len(s1), cap(s1))
	//s1 = append(s1, 1)
	//t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	//
	s2 = append(s2, 2)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 100000; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := months[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := months[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(months)
}

func TestSliceCompare(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//if a == b {
	//	t.Log("equal")
	//}
}

func TestSlicePratice(t *testing.T) {
	var arr = []string{"red", "black", "", "", "pink", "blue"}
	result := removeEmptyStr2(arr)
	fmt.Printf("len:%d,cap:%d\n", len(result), cap(result))
	fmt.Println(result)
}

func removeEmptyStr(data []string) []string {
	result := data[:0]
	fmt.Printf("len:%d,cap:%d\n", len(result), cap(result))
	for _, ele := range data {
		if ele != "" {
			result = append(result, ele)
		}
	}
	return result
}

func removeEmptyStr2(data []string) []string {
	var i int = 0
	for _, ele := range data {
		if ele != "" {
			data[i] = ele
			i++
		}
	}
	return data[:i]
}

/**
去掉重复的字符串
*/
func TestSlicePratice1(t *testing.T) {
	var arr = []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	result := make([]string, 0)
	for _, col := range arr {
		var flag = false
		for _, col1 := range result {
			if col1 == col {
				flag = true
				break
			}
		}
		if !flag {
			result = append(result, col)
		}
	}
	t.Log(result)
}

func TestSliceCopy(t *testing.T) {
	arr := []string{"5", "6", "7", "8", "9"}
	result := sliceCopy(arr)
	t.Log(result)
}

func sliceCopy(data []string) []string {
	len := len(data)
	if len%2 == 0 {
		return data
	}
	half := len / 2
	preHalf := data[:half]
	back := data[1:len]
	copy(back, preHalf)
	return back
}
