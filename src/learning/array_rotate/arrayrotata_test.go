package array_rotate

import (
	"errors"
	"testing"
)

//环状结构  数组长度为n 将原数组下标为i的元素放到新数组下标为(i+k)%n的位置
func rotate1(nums []int, k int) {
	len := len(nums)
	newArr := make([]int, len)
	for i, v := range nums {
		newArr[(i+k)%len] = v
	}
	copy(newArr, nums)
}

func rotate(nums []int, k int) ([]int, error) {
	if k < 0 || nums == nil {
		return nil, errors.New("This is an error")
	}
	len := len(nums)
	if k > len {
		k = k % len
	}
	result := make([]int, len, len)
	for j := k; j > 0; j-- {
		result[k-j] = nums[len-j]
	}
	for i := k; i < len; i++ {
		result[i] = nums[i-k]
	}
	return result, nil
}

func TestArrayRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr2, err := rotate(arr, 5)
	if err == nil {
		for _, v := range arr2 {
			t.Log(v)
		}
	}
}
