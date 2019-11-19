package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programer interface {
	WriteHelloWorld() Code
}

type GoProgramer struct{}

func (p *GoProgramer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgramer struct{}

func (p *JavaProgramer) WriteHelloWorld() Code {
	return "System.Println(\"Hello World\")"
}

func WriteFirstProgram(p Programer) {
	fmt.Printf(" %T %v \n", p, p.WriteHelloWorld())
}

func TestPolyMorphism(t *testing.T) {
	goProgramer := new(GoProgramer)
	javaProgramer := new(JavaProgramer)
	WriteFirstProgram(goProgramer)
	WriteFirstProgram(javaProgramer)
}
