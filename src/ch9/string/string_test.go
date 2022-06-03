package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = 'a'
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"

	c := []rune(s)
	t.Log(len(c))
	t.Logf("%s的unicode is %x", s, c[0])
	t.Logf("%s的utf8 is %x", s, s)

}

func TestStringRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c, c)
	}
}
