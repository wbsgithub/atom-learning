package string

import (
	"strconv"
	"testing"
)

func TestStrConv(t *testing.T) {
	atoi, err := strconv.Atoi("a")
	if err == nil {
		t.Log(atoi)
	} else {
		t.Log(err)
	}

	itoa := strconv.Itoa(1000)
	t.Log("mad" + itoa)
}
