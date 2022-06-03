package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 5}
	arr3 := arr2[1:2]
	t.Log(arr3, len(arr3), cap(arr3))
	t.Log(arr[0], arr[1], arr[2])
	t.Log(arr1, arr2)
	t.Log(arr1[1])
	t.Log(len(arr1), cap(arr1))
	arr4 := append(arr3, 5, 5, 7)
	t.Log(len(arr2), cap(arr2))
	t.Log(arr4, len(arr4), cap(arr4))
	slice := append([]byte("hello "), "world"...)
	t.Log(slice)

}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 {
		t.Log(e)
	}

	c := [2][2]int{{1, 1}, {2, 3}}
	t.Log(c[1])
}

func TestSliceArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	arr2_sec := a[1:3]
	t.Log(arr2_sec)
	t.Log(len(arr2_sec), cap(arr2_sec))
}

func TestArrayMy(t *testing.T) {
	arr8 := [...]int{1, 1}
	t.Log(len(arr8))
	t.Log(cap(arr8))
	sli := arr8[:]
	t.Log(len(sli))
	t.Log(cap(sli))
}
func TestArrayList(t *testing.T) {
	arr8 := [...]int{4, 1, 5, 5}
	for i := 0; i < len(arr8); i++ {
		t.Log(arr8[i])
	}

	arr4 := arr8[1:2]
	t.Log(arr4)
	arr5 := arr8[:]
	t.Log(arr5)
	arr6 := arr8[2:]
	t.Log(arr6)
	arr7 := arr8[4:]
	t.Log(arr7)
	arr9 := arr8[:1]
	t.Log(arr9)

	for index, value := range arr8 {
		t.Log(index, value)
	}

	for _, value := range arr8 {
		t.Log(value)
	}

	a := [2][2]int{{1, 2}, {3, 4}}
	for index, value := range a {
		t.Log(index, value)
		for ind, val := range value {
			t.Log(ind, val)
		}
	}




}
