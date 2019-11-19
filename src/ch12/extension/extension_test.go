package extension

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	fmt.Println("...", host)
}

type Dog struct {
	Pet
}

func (p *Dog) Speak() {
	fmt.Print("Wang")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("CAO")
}
