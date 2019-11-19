package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		temp := a
		a = b
		b = temp + a
	}
	t.Log()
}
