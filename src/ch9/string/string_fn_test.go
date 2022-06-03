package string_test

import (
	"strings"
	"testing"
)

func Test(t *testing.T){
	s := ""
	split := strings.Split(s, "")
	t.Log(len(split))
	t.Log(cap(split))
	for _,c:=range split{
		t.Log(c)
	}
}