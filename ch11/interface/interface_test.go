package interface_test

import (
	"fmt"
	"testing"
)

type Programer interface {
	WriteHelloWorld() string
}

type GoProgramer struct {
}

func (g *GoProgramer) WriteHelloWorld() string {
	return fmt.Sprint("Hello World!")
}

func TestClient(t *testing.T) {
	var p Programer
	p = new(GoProgramer)
	t.Log(p.WriteHelloWorld())
}
