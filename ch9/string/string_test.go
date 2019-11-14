package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"

	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	/// range拿出来的c,自动转为了rune的Unicode
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
